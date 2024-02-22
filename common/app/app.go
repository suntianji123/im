package app

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/common/acceptor"
	"github.com/im/common/data"
)

type App struct {
	builder    *Builder
	tcpServer  *acceptor.TcpAcceptor
	httpServer *acceptor.HttpAcceptor
	data       *data.Data
}

func (p *App) Start() error {
	//启动数据服务
	if p.data != nil {
		defer func() {
			err := p.data.Close()
			if err != nil {
				logger.Errorf("App close data failed:%v", err)
			}
		}()

		err := p.data.Init()
		if err != nil {
			logger.Errorf("App Init data failed:%v", err)
			return err
		}
	}

	//启动http服务
	if p.httpServer != nil {
		err := p.httpServer.Init()
		if err != nil {
			logger.Errorf("App start httpServer failed:%v", err)
			return err
		}
	}

	//启动tcp服务
	if p.tcpServer != nil {
		defer p.tcpServer.Stop()
		p.tcpServer.Init()
	}

	select {}
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
