package agent

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/im/common/codec"
	"github.com/im/common/constants"
	"github.com/im/common/session"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type Agent struct {
	conn             net.Conn
	chDie            chan struct{}
	chStopHeartBeat  chan struct{}
	chStopWrite      chan struct{}
	chSend           chan pendingWrite
	closeMutex       sync.Mutex
	session          *session.Session
	sessionPool      *session.SessionPool
	heartBeatTimeout time.Duration
	lastAt           int64
	state            int32
}

type pendingWrite struct {
	data []byte
}

type AgentFactory struct {
	sessionPool *session.SessionPool
}

func NewAgentFactory() *AgentFactory {
	return &AgentFactory{
		sessionPool: session.NewSessionPool(),
	}
}

func (p *AgentFactory) CreateAgent(conn net.Conn) *Agent {
	return p.newAgent(conn)
}

func (p *AgentFactory) newAgent(conn net.Conn) *Agent {
	a := &Agent{
		conn:             conn,
		chDie:            make(chan struct{}),
		heartBeatTimeout: time.Duration(1) * time.Minute,
		chStopHeartBeat:  make(chan struct{}),
		chStopWrite:      make(chan struct{}),
		chSend:           make(chan pendingWrite),

		sessionPool: p.sessionPool,
	}

	s := p.sessionPool.NewSession(a)
	a.session = s
	return a
}

func (a *Agent) Handle() {
	defer func() {
		a.Close()
	}()

	go a.heartBeat()
	go a.write()
	<-a.chDie
}

func (a *Agent) write() {
	defer a.Close()

	for {
		select {
		case write := <-a.chSend:
			bytes := codec.IntToBytes(4 + len(write.data))
			bytes = append(bytes, write.data...)
			if _, err := a.conn.Write(bytes); err != nil {
				logger.Errorf("Agent:%s Write failed:%v", a.conn.RemoteAddr().String(), err)
				return
			}
			//logger.Infof("send message:%s success", string(bytes))
		case <-a.chStopWrite:
			return
		}
	}
}

func (a *Agent) heartBeat() {
	ticker := time.NewTicker(a.heartBeatTimeout)
	defer func() {
		ticker.Stop()
		logger.Warnf("agent heartbeart close agent")
		a.Close()
	}()

	for {
		select {
		case <-ticker.C:
			deadline := time.Now().Add(-2 * a.heartBeatTimeout).UnixMilli()
			if atomic.LoadInt64(&a.lastAt) < deadline {
				logger.Warnf("Agent:%s heartBeat timeout", a.conn.RemoteAddr())
				return
			}

			select {
			case <-a.chDie:
				return
			case <-a.chStopHeartBeat:
				return
			default:

			}

		case <-a.chDie:
			return
		case <-a.chStopHeartBeat:
			return
		}
	}
}

func (a *Agent) Send(ctx context.Context, msg proto.Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		logger.Errorf("Agent Send failed:%v", err)
		return err
	}

	select {
	case a.chSend <- pendingWrite{
		data: data,
	}:
	case <-a.chDie:
	default:
	}
	return nil
}

func (a *Agent) Close() error {
	a.closeMutex.Lock()
	defer a.closeMutex.Unlock()

	if a.GetStatus() == constants.StatusClosed {
		//logger.Errorf("close closed agent:%s", a.conn.RemoteAddr().String())
		return constants.ErrCloseClosedSession
	}
	a.SetStatus(constants.StatusClosed)
	logger.Infof("close agent remote address:%s", a.conn.RemoteAddr().String())

	select {
	case <-a.chDie:
	default:
		close(a.chStopWrite)
		close(a.chStopHeartBeat)
		close(a.chDie)
	}

	return a.conn.Close()
}

func (a *Agent) GetStatus() int32 {
	return atomic.LoadInt32(&a.state)
}

func (a *Agent) SetStatus(state int32) {
	atomic.StoreInt32(&a.state, state)
}

func (a *Agent) GetAddr() string {
	return a.conn.RemoteAddr().String()
}

func (a *Agent) SetLastAt() {
	atomic.StoreInt64(&a.lastAt, time.Now().UnixMilli())
}

func (a *Agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *Agent) GetSession() *session.Session {
	return a.session
}

func (a *Agent) IsActive() bool {
	return atomic.LoadInt32(&a.state) != constants.StatusClosed
}

func (a *AgentFactory) GetSessionPool() *session.SessionPool {
	return a.sessionPool
}
