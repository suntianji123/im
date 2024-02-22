package rpc

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/im/common/api"
)

var UserServiceClientImpl = &api.UserServiceClientImpl{}
var FriendServiceClientImpl = &api.FriendServiceClientImpl{}
var AuthServiceClientImpl = &api.AuthServiceClientImpl{}
var SyncServiceClientImpl = &api.SyncServiceClientImpl{}

func init() {
	config.SetConsumerService(UserServiceClientImpl)
	config.SetConsumerService(FriendServiceClientImpl)
	config.SetConsumerService(AuthServiceClientImpl)
	config.SetConsumerService(SyncServiceClientImpl)
}
