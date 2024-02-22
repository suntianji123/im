package main

import (
	cp "github.com/im/common/app"
	"github.com/im/signal/auth/app"
	_ "github.com/im/signal/auth/pkg/service"
)

func main() {
	//注册rpc client
	app.App = cp.NewBuilder().Build()
	err := app.App.Start()
	if err != nil {
		panic(err)
	}

}
