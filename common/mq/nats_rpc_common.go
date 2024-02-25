package mq

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/nats-io/nats.go"
)

func setupNatsConn(connectString string, appDie chan bool, options ...nats.Option) (*nats.Conn, error) {
	natsOptions := append(options, nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
		if err != nil {
			logger.Errorf("Nats connect to %s err:%v", connectString, err)
		}

	}), nats.ReconnectHandler(func(conn *nats.Conn) {
		logger.Infof("Nats reconnect to %s", connectString)
	}), nats.ClosedHandler(func(conn *nats.Conn) {
		err := conn.LastError()
		if err == nil {
			logger.Warnf("Nats closed with no error")
			return
		}
		logger.Errorf("Nats closed with error:%v", err)
		if appDie != nil {
			appDie <- true
		}
	}), nats.ErrorHandler(func(conn *nats.Conn, sub *nats.Subscription, err error) {
		if err == nats.ErrSlowConsumer {
			dropped, _ := sub.Dropped()
			logger.Warnf("nats slow consumer on Subject %q: dropped %d messages\n",
				sub.Subject, dropped)
		} else {
			logger.Errorf(err.Error())
		}
	}))

	nc, err := nats.Connect(connectString, natsOptions...)
	if err != nil {
		logger.Errorf("setupNatsConn failed:%v", err)
		return nil, err
	}
	return nc, nil
}
