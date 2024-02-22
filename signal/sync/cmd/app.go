package main

import (
	_ "github.com/im/signal/sync/pkg/service"

	cp "github.com/im/common/app"
	"github.com/im/signal/sync/app"

	_ "github.com/im/signal/sync/rpc"
)

func main() {
	app.App = cp.NewBuilder().Build()
	if err := app.App.Start(); err != nil {
		panic(err)
	}
}
