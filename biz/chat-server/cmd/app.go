package main

import (
	"github.com/im/biz/chat-server/app"
	_ "github.com/im/biz/chat-server/handler"
	_ "github.com/im/biz/chat-server/pkg/service"
	_ "github.com/im/biz/chat-server/rpc"
	cp "github.com/im/common/app"
)

func main() {
	app.App = cp.NewBuilder().Build()
	app.App.RegisterIdGenerator()
	if err := app.App.Start(); err != nil {
		panic(err)
	}
}
