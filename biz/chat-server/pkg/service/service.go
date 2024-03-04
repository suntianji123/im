package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/golang/protobuf/proto"
)

type Handler interface {
	Handle(ctx context.Context, message proto.Message) error
	Message() proto.Message
}

var Handlers = make(map[int]Handler)

func init() {
	config.SetProviderService(NewChatServiceServerImpl())
}
