package main

import (
	_ "github.com/im/signal/sync/pkg/service"

	cp "github.com/im/common/app"
	"github.com/im/signal/sync/app"

	m "github.com/im/common/mq"
	"github.com/im/signal/sync/mq"
	_ "github.com/im/signal/sync/rpc"
)

func main() {
	app.App = cp.NewBuilder().Build()
	app.App.GetNatsServer().RegisterSubs([]m.SubHandler{
		mq.NewSyncSubHandler(),
	})
	if err := app.App.Start(); err != nil {
		panic(err)
	}
}
