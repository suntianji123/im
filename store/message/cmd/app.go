package main

import (
	cp "github.com/im/common/app"
	"github.com/im/store/message/app"
	_ "github.com/im/store/message/pkg/service"
)

func main() {

	app.App = cp.NewBuilder().Build()
	if err := app.App.Start(); err != nil {
		panic(err)
	}

}
