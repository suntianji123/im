package handler

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/im/common/api"
	"github.com/im/common/session"
	"github.com/im/signal/connector/rpc"
)

type pChatMsgHandler struct{}

func (p *pChatMsgHandler) Message() proto.Message {
	return &api.PTransUp{}
}

func (p *pChatMsgHandler) Handler(ctx context.Context, session *session.Session, msg proto.Message) error {
	req := msg.(*api.PTransUp)
	routeKey := req.RouteKey
	if len(routeKey) == 0 {
		routeKey = fmt.Sprintf("%d", req.Uid)
	}
	_, err := rpc.ChatServiceClientImpl.Handle(ctx, &api.HandleReq{Key: routeKey, Data: req.Data})
	if err != nil {
		logger.Errorf("pChatMsgHandler Handler %v failed:%v", req, err)
		return err
	}
	return nil
}
