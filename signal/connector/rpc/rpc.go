package rpc

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/im/common/api"
)

var AuthServiceClientImpl = &api.AuthServiceClientImpl{}
var ChatServiceClientImpl = &api.ChatServiceClientImpl{}

func init() {
	config.SetConsumerService(AuthServiceClientImpl)
	config.SetConsumerService(ChatServiceClientImpl)
}
