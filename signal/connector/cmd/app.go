package main

import (
	cp "github.com/im/common/app"
	"github.com/im/signal/connector/app"
	"github.com/im/signal/connector/handler"
	_ "github.com/im/signal/connector/rpc"
)

func main() {
	app.App = cp.NewBuilder().Build()

	app.App.GetTcpAcceptor().RegisterHandlers(handler.Handlers)

	if err := app.App.Start(); err != nil {
		panic(err)
	}
}
