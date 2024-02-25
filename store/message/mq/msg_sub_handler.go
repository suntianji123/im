package mq

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/golang/protobuf/proto"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/data"
	"github.com/im/common/data/ent"
	"github.com/im/common/mq"
	"github.com/im/common/util"
	"github.com/nats-io/nats.go"
)

type msgSubHandler struct {
	mq.BaseSubHandler
}

func NewMsgSubHandler() *msgSubHandler {
	return &msgSubHandler{
		mq.BaseSubHandler{
			Subject: constants.MqMessageSubject,
			MsgChan: make(chan *nats.Msg),
		},
	}
}

func (p *msgSubHandler) Handle(msg *nats.Msg) error {
	req := &api.PChatMsgSendReq{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		logger.Errorf("MsgSubHandler handle proto unmarshal failed:%v", err)
		return err
	}

	body, err := util.ProtocalUtil.Serialize(req)
	if err != nil {
		logger.Errorf("MsgSubHandler handle ProtocalUtil Serailize failed:%v", err)
		return err
	}
	ctx := context.Background()
	data.MsgBodyRepo.Create(ctx, &ent.MsgBody{MsgID: req.MsgId, Body: body, Cts: req.Cts})

	data.ImMsgRepo.Create(ctx, &ent.IMMsg{
		Sid:       util.ProtocalUtil.GetChatSid(req),
		FromUID:   req.FromUid,
		FromAppid: int(req.AppId),
		ToUID:     req.ToUid,
		ToAppid:   int(req.ToAppId),
		Channel:   int(req.Channel),
		MsgID:     req.MsgId,
		Cts:       req.Cts,
	})
	return nil
}
