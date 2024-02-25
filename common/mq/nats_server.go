package mq

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/common/conf"
	"github.com/nats-io/nats.go"
	"time"
)

type SubHandler interface {
	GetSubject() string
	Handle(msg *nats.Msg) error
	GetMessageChan() chan *nats.Msg
	SetSub(sub *nats.Subscription)
	Close() error
}

type BaseSubHandler struct {
	Subject string
	Sub     *nats.Subscription
	MsgChan chan *nats.Msg
}

type NatsServer struct {
	conn        *nats.Conn
	running     bool
	dieChan     chan bool
	config      *conf.NatsConfig
	subHandlers []SubHandler
}

func NewNatsServer(config *conf.NatsConfig) *NatsServer {
	return &NatsServer{
		config:  config,
		dieChan: make(chan bool),
	}
}

func (ns *NatsServer) RegisterSubs(subs []SubHandler) {
	ns.subHandlers = subs
}

func (ns *NatsServer) Init() error {
	logger.Infof("Nats connect to %s with timeout:%d", ns.config.Addr, ns.config.ConnectTimeout)
	conn, err := setupNatsConn(ns.config.Addr, ns.dieChan,
		nats.Timeout(time.Duration(ns.config.ConnectTimeout)*time.Second),
		nats.MaxReconnects(ns.config.MaxReconnectionRetries))
	if err != nil {
		logger.Errorf("NatsServer init setupNatsConn failed:%v", err)
		return err
	}

	ns.running = true
	ns.conn = conn

	for _, h := range ns.subHandlers {
		sub, err := ns.conn.ChanSubscribe(h.GetSubject(), h.GetMessageChan())
		if err != nil {
			logger.Errorf("NatsServer subscribe Subject:%s failed:%v", h.GetSubject(), err)
			return err
		}
		h.SetSub(sub)
	}

	for _, h := range ns.subHandlers {
		go func() {
			for {
				select {
				case msg := <-h.GetMessageChan():
					err = h.Handle(msg)
					if err != nil {
						logger.Errorf("nats server subscribe Subject:%s,msg:%v,Handler failed:%v", h.GetSubject(), msg, err)
					}
				case <-ns.dieChan:
					ns.Stop()
					return
				}
			}
		}()
	}

	return nil
}

func (ns *NatsServer) Stop() error {
	ns.running = false
	for _, h := range ns.subHandlers {
		err := h.Close()
		if err != nil {
			logger.Errorf("natServer subHandler subject:%s close failed:%v", h.GetSubject(), err)
		}
	}
	//close(ns.dieChan)
	ns.conn.Close()
	logger.Warnf("Nat server closed....")
	return nil
}

func (p *BaseSubHandler) GetSubject() string {
	return p.Subject
}

func (p *BaseSubHandler) Handle(msg *nats.Msg) error {
	return nil
}

func (p *BaseSubHandler) GetMessageChan() chan *nats.Msg {
	return p.MsgChan
}

func (p *BaseSubHandler) SetSub(sub *nats.Subscription) {
	p.Sub = sub
}

func (p *BaseSubHandler) Close() error {
	//close(p.MsgChan)
	//err := p.Sub.Unsubscribe()
	//if err != nil {
	//	logger.Errorf("BaseSubHandler subject:%s close failed:%v", p.Subject, err)
	//	//return err
	//}
	return nil
}
