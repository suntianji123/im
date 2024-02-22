package main

import (
	"github.com/im/account/app"
	_ "github.com/im/account/pkg/service"
	_ "github.com/im/account/rpc"
	cp "github.com/im/common/app"
)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func main() {
	app.App = cp.NewBuilder().Build()
	if err := app.App.Start(); err != nil {
		panic(err)
	}
}
