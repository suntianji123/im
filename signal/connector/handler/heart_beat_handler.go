package handler

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/golang/protobuf/proto"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/session"
	"time"
)

type heartBeatHandler struct{}

func (p *heartBeatHandler) Message() proto.Message {
	return &api.PPing{}
}

func (p *heartBeatHandler) Handler(ctx context.Context, session *session.Session, msg proto.Message) error {
	ping := msg.(*api.PPing)
	err := session.Push(ctx, &api.PPong{
		Uri:   int32(constants.Pong),
		Uid:   ping.Uid,
		AppId: ping.AppId,
		Cts:   ping.Cts,
		Sts:   time.Now().UnixMilli(),
	})

	if err != nil {
		logger.Errorf("handle pingMsg failed:%v", err)
		return err
	}
	//logger.Infof("发送pong消息给agent:%s成功", agent.GetAddr())
	return nil
}
