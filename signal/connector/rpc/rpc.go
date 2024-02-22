package rpc

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/im/common/api"
)

var AuthServiceClientImpl = &api.AuthServiceClientImpl{}

func init() {
	config.SetConsumerService(AuthServiceClientImpl)
}
