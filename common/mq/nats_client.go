package mq

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/common/conf"
	"github.com/im/common/constants"
	"github.com/nats-io/nats.go"
	"time"
)

type NatsClient struct {
	conn       *nats.Conn
	running    bool
	appDieChan chan bool
	config     *conf.NatsConfig
}

func NewNatsClient(config *conf.NatsConfig) *NatsClient {
	return &NatsClient{
		config:     config,
		appDieChan: make(chan bool),
	}
}

func (ns *NatsClient) Init() error {

	logger.Infof("NatsClient connect to %s with timeout:%d", ns.config.Addr, ns.config.ConnectTimeout)
	conn, err := setupNatsConn(ns.config.Addr,
		ns.appDieChan,
		nats.MaxReconnects(ns.config.MaxReconnectionRetries),
		nats.Timeout(time.Duration(ns.config.ConnectTimeout)*time.Second))
	if err != nil {
		logger.Errorf("NatsClient Init setupNatsConn failed:%v", err)
		return err
	}
	ns.conn = conn
	ns.running = true

	go func() {
		select {
		case <-ns.appDieChan:
			ns.Stop()
			return
		}
	}()
	return nil
}

func (ns *NatsClient) Publish(subject string, data []byte) error {
	if !ns.running {
		return constants.ErrNatsNotRunning
	}

	err := ns.conn.Publish(subject, data)
	if err != nil {
		logger.Errorf("NatsClient Publish Subject:%s failed:%v", err)
		return err
	}
	return nil
}

func (ns *NatsClient) Stop() error {
	ns.running = false
	ns.conn.Close()
	logger.Warnf("Nats client closed....")
	return nil
}
