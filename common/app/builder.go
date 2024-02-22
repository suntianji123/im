package app

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/im/common/acceptor"
	"github.com/im/common/conf"
	"github.com/im/common/data"
	"github.com/mitchellh/mapstructure"
)

type Builder struct {
	dataConfig       *conf.DataConfig
	httpServerConfig *conf.HttpServerConfig
	tcpServerConfig  *conf.TcpServerConfig
}

func NewBuilder() *Builder {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	builder := &Builder{}
	if v, ok := config.GetRootConfig().Custom.ConfigMap["Data"]; ok {
		dataConfig := &conf.DataConfig{
			DbConfig:    &conf.DatabaseConfig{},
			RedisConfig: &conf.RedisConfig{},
		}
		err = mapstructure.Decode(v, &dataConfig)
		if err != nil {
			panic(err)
		}
		builder.dataConfig = dataConfig
	}

	if v, ok := config.GetRootConfig().Custom.ConfigMap["Http"]; ok {
		httpConfig := &conf.HttpServerConfig{}
		err = mapstructure.Decode(v, &httpConfig)
		if err != nil {
			panic(err)
		}
		builder.httpServerConfig = httpConfig
	}

	if v, ok := config.GetRootConfig().Custom.ConfigMap["Tcp"]; ok {
		tcpConfig := &conf.TcpServerConfig{}
		err = mapstructure.Decode(v, &tcpConfig)
		if err != nil {
			panic(err)
		}
		builder.tcpServerConfig = tcpConfig
	}

	return builder
}

func (p *Builder) Build() *App {
	var d *data.Data
	if p.dataConfig != nil {
		d = data.NewData(p.dataConfig)
	}

	var httpServer *acceptor.HttpAcceptor
	if p.httpServerConfig != nil {
		httpServer = acceptor.NewHttpAcceptor(p.httpServerConfig.Port)
	}

	var tcpServer *acceptor.TcpAcceptor
	if p.tcpServerConfig != nil {
		tcpServer = acceptor.NewTcpAcceptor(p.tcpServerConfig.Port)
	}

	return &App{
		builder:    p,
		data:       d,
		httpServer: httpServer,
		tcpServer:  tcpServer,
	}
}
