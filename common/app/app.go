package app

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/im/common/acceptor"
	"github.com/im/common/data"
	"github.com/im/common/mq"
	"github.com/im/common/util"
	"github.com/im/common/uuid"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	builder     *Builder
	tcpServer   *acceptor.TcpAcceptor
	httpServer  *acceptor.HttpAcceptor
	data        *data.Data
	natsClient  *mq.NatsClient
	natsServer  *mq.NatsServer
	dieChan     chan bool
	idGenerator *uuid.IdGenerator
}

func (p *App) Start() error {
	if p.builder.channelConfig != nil {
		util.ChannelUtil.Init(p.builder.channelConfig)
		logger.Infof("App Init ChannelUtil success...")
	}

	if p.idGenerator != nil {
		err := p.idGenerator.Init()
		if err != nil {
			logger.Errorf("App Init IdGenerator failed:%v", err)
			return err
		}
	}

	//启动数据服务
	if p.data != nil {
		err := p.data.Init()
		if err != nil {
			logger.Errorf("App Init data failed:%v", err)
			return err
		}
		logger.Infof("App start data success...")
	}

	//启动http服务
	if p.httpServer != nil {
		err := p.httpServer.Init()
		if err != nil {
			logger.Errorf("App start httpServer failed:%v", err)
			return err
		}
		logger.Infof("App start httpServer success...")
	}

	//启动tcp服务
	if p.tcpServer != nil {
		p.tcpServer.Init()
		logger.Infof("App start tcpServer success...")
	}

	if p.natsServer != nil {
		err := p.natsServer.Init()
		if err != nil {
			logger.Errorf("App start natsServer failed:%v", err)
			return err
		}
		logger.Infof("App start natsServer sucess...")
	}

	if p.natsClient != nil {
		err := p.natsClient.Init()
		if err != nil {
			logger.Errorf("App start natsClient failed:%v", err)
			return err
		}
		logger.Infof("App start natsClient success")
	}

	sg := make(chan os.Signal)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)
	select {
	case <-p.dieChan:
		logger.Warnf("The app will down in a few time")
	case s := <-sg:
		logger.Warnf("got signal:%v,shutdown...", s)
		close(p.dieChan)
	}

	p.Stop()
	return nil
}

func (p *App) GetHttpAcceptor() *acceptor.HttpAcceptor {
	return p.httpServer
}

func (p *App) GetTcpAcceptor() *acceptor.TcpAcceptor {
	return p.tcpServer
}

func (p *App) GetData() *data.Data {
	return p.data
}

func (p *App) Stop() {
	logger.Infof("close other service...")
	var err error
	if p.httpServer != nil {
		err = p.httpServer.Stop()
		if err != nil {
			logger.Errorf("App close httpServer with error:%v", err)
		}

	}

	if p.tcpServer != nil {
		err = p.tcpServer.Stop()
		if err != nil {
			logger.Errorf("App close tcpServer with error:%v", err)
		}

	}

	if p.natsClient != nil {
		err = p.natsClient.Stop()
		if err != nil {
			logger.Errorf("App close natsClient with error:%v", err)
		}
	}

	if p.natsServer != nil {
		err = p.natsServer.Stop()
		if err != nil {
			logger.Errorf("App close natsServer with error:%v", err)
		}
	}

	if p.data != nil {
		err = p.data.Close()
		if err != nil {
			logger.Errorf("App close data with error:%v", err)
		}
	}

	if p.idGenerator != nil {
		p.idGenerator.Stop()
		logger.Warnf("App close IdGenerator success...")
	}

	if err != nil {
		logger.Warnf("App close other service success....")
	}
}

func (p *App) GetNatsClient() *mq.NatsClient {
	return p.natsClient
}

func (p *App) GetNatsServer() *mq.NatsServer {
	return p.natsServer
}

func (p *App) RegisterIdGenerator() {
	p.idGenerator = uuid.NewIdGenerator()
}

func (p *App) GetIdGenerator() *uuid.IdGenerator {
	return p.idGenerator
}
