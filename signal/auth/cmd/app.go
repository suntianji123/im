package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/im/signal/auth/pkg/service"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	select {}
}
