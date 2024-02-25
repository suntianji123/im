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
	"github.com/im/common/uuid"
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

	me := &ent.ChatList{
		UID:      req.FromUid,
		ChatID:   req.ToUid,
		Channel:  int(req.Channel),
		MaxMsgID: req.MsgId,
		Uts:      uuid.GetTs(req.MsgId),
		Type:     constants.ChatTypeChat,
	}

	peer := &ent.ChatList{
		UID:      req.ToUid,
		ChatID:   req.FromUid,
		Channel:  int(req.Channel),
		MaxMsgID: req.MsgId,
		Uts:      uuid.GetTs(req.MsgId),
		Type:     constants.ChatTypeChat,
	}

	data.ChatListRepo.CreateOrUpdate(context.Background(), me, peer)
	return nil
}
