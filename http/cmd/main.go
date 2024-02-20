package main

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/im/http/router"
	_ "github.com/im/http/router/handler"
	_ "github.com/im/http/rpc"
)

func main() {
	//注册rpc client
	err := config.Load()
	if err != nil {
		logger.Errorf("启动http出错:%v", err)
	}
	//监听http请求
	err = router.NewRouter().Init()
	if err != nil {
		panic(err)
	}

}
