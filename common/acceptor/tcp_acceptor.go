package acceptor

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/im/common/agent"
	"github.com/im/common/codec"
	"github.com/im/common/constants"
	"github.com/im/common/session"
	"io"
	"net"
)

type TcpAcceptor struct {
	addr         int
	connChan     chan *UserChan
	listener     net.Listener
	running      bool
	msgChan      chan *packet
	handlers     map[int]TcpHandler
	agentFactory *agent.AgentFactory
}

type TcpHandler interface {
	Handler(context.Context, *session.Session, proto.Message) error
	Message() proto.Message
}

type UserChan struct {
	conn       net.Conn
	remoteAddr net.Addr
}

type packet struct {
	uri     int
	ctx     context.Context
	session *session.Session
	message proto.Message
}

func NewTcpAcceptor(addr int) *TcpAcceptor {
	return &TcpAcceptor{
		addr:         addr,
		connChan:     make(chan *UserChan),
		msgChan:      make(chan *packet),
		running:      false,
		agentFactory: agent.NewAgentFactory(),
	}
}

func (a *TcpAcceptor) RegisterHandlers(handlers map[int]TcpHandler) {
	a.handlers = handlers
}

func (a *TcpAcceptor) Init() {
	go a.Dispatcher()

	go func() {
		for acc := range a.connChan {
			go a.Handler(acc)
		}
	}()

	go a.ListenAndServe()
}

func (a *TcpAcceptor) Stop() error {
	a.running = false
	err := a.listener.Close()
	if err != nil {
		logger.Errorf("close tcp server failed:%v", err)
		return err
	} else {
		logger.Warnf("close tcp server success")
		return nil
	}
}

func (a *TcpAcceptor) ListenAndServe() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.addr))
	if err != nil {
		logger.Errorf("TcpAcceptor ListenAndServe failed:%v", err)
	}

	a.listener = listener
	a.running = true
	a.serve()
}

func (a *TcpAcceptor) serve() {
	defer close(a.connChan)
	for a.running {
		conn, err := a.listener.Accept()
		if err != nil {
			logger.Errorf("Failed to accept Tcp:%v", err)
			continue
		}
		logger.Infof("accept Tcp:%s", conn.RemoteAddr().String())

		a.connChan <- &UserChan{
			conn:       conn,
			remoteAddr: conn.RemoteAddr(),
		}
	}

}

func (a *TcpAcceptor) Dispatcher() {
	defer func() {
		close(a.msgChan)

	}()

	select {
	case p := <-a.msgChan:
		err := a.handlers[p.uri].Handler(p.ctx, p.session, p.message)
		if err != nil {
			logger.Errorf("TcpAcceptor Dispatcher failed:%v", err)
		}
	}

}

func (a *TcpAcceptor) GetUserChan() chan *UserChan {
	return a.connChan
}

func (a *TcpAcceptor) Handler(conn *UserChan) {
	agent := a.agentFactory.CreateAgent(conn.conn)
	defer func() {
		agent.GetSession().Close()
		logger.Infof("TcpAcceptor close session:%s", agent.GetSession().RemoteAddr().String())
	}()
	go agent.Handle()

	for {
		bytes, err := conn.GetNextMessage()
		if err != nil {
			logger.Errorf("Read conn message failed:%v", err)
			return
		}

		m := make(map[string]interface{})
		err = json.Unmarshal(bytes, &m)
		if err != nil {
			logger.Errorf("TcpAcceptor Handler json umarshal failed:%v", err)
			return
		}

		uri := int(m["uri"].(float64))
		h, ok := a.handlers[uri]
		if !ok {
			logger.Errorf("TcpAcceptor Handler json umarshal failed:未注册uri:%d的处理事件", uri)
			continue
		}

		msg := h.Message()
		err = jsonpb.UnmarshalString(string(bytes), msg)
		if err != nil {
			logger.Errorf("TcpAcceptor Handler json umarshal failed:%v", err)
		}
		select {
		case a.msgChan <- &packet{
			uri:     uri,
			ctx:     context.Background(),
			session: agent.GetSession(),
			message: msg,
		}:
		default:
		}
		agent.SetLastAt()
	}
}

func (t *UserChan) GetNextMessage() (b []byte, err error) {
	bytes, err := io.ReadAll(io.LimitReader(t.conn, constants.PacketLength))
	if err != nil {
		logger.Errorf("UserChann GetNextMessage failed:%v", err)
		return nil, err
	}

	if len(bytes) == 0 {
		return nil, constants.ErrConnectionClosed
	}

	if len(bytes) < constants.PacketLength {
		return nil, constants.ErrReceivedMsgSmallerThanExpected
	}

	length := codec.BytesToInt(bytes) - constants.PacketLength
	bytes, err = io.ReadAll(io.LimitReader(t.conn, int64(length)))
	if err != nil {
		logger.Errorf("UserChann GetNextMessage failed:%v", err)
		return nil, err
	}

	if len(bytes) < length {
		return nil, constants.ErrReceivedMsgSmallerThanExpected
	}

	return bytes, nil
}
