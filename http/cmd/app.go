package main

import (
	cp "github.com/im/common/app"
	"github.com/im/http/app"
	"github.com/im/http/handler"
	"github.com/im/http/interceptor"
	_ "github.com/im/http/rpc"
)

func main() {
	//注册rpc client
	app.App = cp.NewBuilder().Build()
	app.App.GetHttpAcceptor().RegisterHandlers(handler.Handlers)
	app.App.GetHttpAcceptor().RegisterInterceptor(interceptor.AuthIntecetor)

	err := app.App.Start()
	if err != nil {
		panic(err)
	}

}
