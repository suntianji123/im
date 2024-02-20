package main

import (
	_ "github.com/im/account/pkg/service"

	"dubbo.apache.org/dubbo-go/v3/config"

	_ "dubbo.apache.org/dubbo-go/v3/imports"

	_ "github.com/im/account/rpc"

	"github.com/im/common/data"
)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	//初始化数据库 redis
	err := data.Data.Init()
	if err != nil {
		panic(err)
	}

	select {}
}
