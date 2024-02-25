package rpc

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/im/common/api"
)

var MessageServiceClientImpl = &api.MessageServiceClientImpl{}

func init() {
	config.SetConsumerService(MessageServiceClientImpl)
}
