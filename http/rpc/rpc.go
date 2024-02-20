package rpc

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/im/common/api"
)

var UserServiceClientImpl = &api.UserServiceClientImpl{}
var FriendServiceClientImpl = &api.FriendServiceClientImpl{}
var AuthServiceClientImpl = &api.AuthServiceClientImpl{}

func init() {
	config.SetConsumerService(UserServiceClientImpl)
	config.SetConsumerService(FriendServiceClientImpl)
	config.SetConsumerService(AuthServiceClientImpl)
}
