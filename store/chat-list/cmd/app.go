package main

import (
	cp "github.com/im/common/app"
	m "github.com/im/common/mq"
	"github.com/im/store/chat-list/app"
	"github.com/im/store/chat-list/mq"
)

func main() {

	app.App = cp.NewBuilder().Build()

	//消费mq
	app.App.GetNatsServer().RegisterSubs([]m.SubHandler{
		mq.NewMsgSubHandler(),
	})

	if err := app.App.Start(); err != nil {
		panic(err)
	}
}
