package handler

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/golang/protobuf/proto"
	"github.com/im/biz/chat-server/app"
	"github.com/im/biz/chat-server/pkg/service"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/util"
	"time"
)

type pChatMsgHandler struct {
}

func (p *pChatMsgHandler) Message() proto.Message {
	return &api.PChatMsgSendReq{}
}

func (p *pChatMsgHandler) Handle(ctx context.Context, message proto.Message) error {
	req := message.(*api.PChatMsgSendReq)
	if req.MsgId <= 0 {
		msgId, err := service.SeqIdService.GetMsgId(req)
		if err != nil {
			logger.Errorf("pChatMsgHandler Handle seqIdService getMsgId failed:%v", err)
			return err
		}
		req.MsgId = msgId
	}

	seqId, err := service.SeqIdService.GetSeqId(ctx, req)
	if err != nil {
		logger.Errorf("pChatMsgHandler Handler seqIdService getSeqId failed:%v", err)
		return err
	}
	req.SeqId = seqId
	req.Sts = time.Now().UnixMilli()

	err = service.MsgService.SaveBody(ctx, req)
	if err != nil {
		logger.Errorf("pChatMsgHandler Handler MsgService SaveBody failed:%v", err)
		return err
	}

	//生产消息
	//写入数据库 chatlist更新会话列表 msg更新msgbody、immsg
	bytes, err := proto.Marshal(req)
	if err != nil {
		logger.Errorf("pChatMsgHandler Handler proto Marshal failed:%v", err)
		return err
	}
	err = app.App.GetNatsClient().Publish(constants.MqMessageSubject, bytes)
	if err != nil {
		logger.Errorf("pChatMsgHandler Handler NatsClient publish failed:%v", err)
		return err
	}

	//同步给自己 同步给他人
	err = p.syncToMe(req)
	if err != nil {
		logger.Errorf("pChatMsgHandler Handler syncToMe data:%v failed:%v", req, err)
		return err
	}

	err = p.syncToOther(req)
	if err != nil {
		logger.Errorf("pChatMsgHandler Handler syncToOther data:%v failed:%v", req, err)
		return err
	}

	return nil
}

func (p *pChatMsgHandler) syncToMe(msg *api.PChatMsgSendReq) error {
	data, err := util.ProtocalUtil.Serialize(util.ProtocalUtil.GetPChatMsgSendReqJson(msg))
	if err != nil {
		logger.Errorf("PChatMsgHandler syncToMe protocalUtil Serialize %v failed:%v", msg, err)
		return err
	}
	req := &api.PSyncReq{
		Uri:      constants.SyncReq,
		AppId:    msg.AppId,
		Uid:      msg.FromUid,
		Channel:  msg.Channel,
		DeviceId: "",
		ChatType: int32(constants.ChatTypeChat),
		MsgId:    msg.MsgId,
		Data:     data,
		SyncPos:  constants.PSyncReqNeedSync,
		PushType: constants.PushTypeNone,
	}

	bytes, err := proto.Marshal(req)
	if err != nil {
		logger.Errorf("PChatMsgHandler syncToMe proto marshal %v failed:%v", req, err)
		return err
	}

	err = app.App.GetNatsClient().Publish(constants.MqSyncSubject, bytes)
	if err != nil {
		logger.Errorf("PChatMsgHandler syncToMe natsClient publish %v failed:%v", req, err)
		return err
	}
	return nil
}

func (p *pChatMsgHandler) syncToOther(msg *api.PChatMsgSendReq) error {
	data, err := util.ProtocalUtil.Serialize(util.ProtocalUtil.GetPChatMsgSendReqJson(msg))
	if err != nil {
		logger.Errorf("PChatMsgHandler syncToOther protocalUtil Serialize %v failed:%v", msg, err)
		return err
	}
	req := &api.PSyncReq{
		Uri:      constants.SyncReq,
		AppId:    msg.AppId,
		Uid:      msg.ToUid,
		Channel:  msg.Channel,
		DeviceId: "",
		ChatType: int32(constants.ChatTypeChat),
		MsgId:    msg.MsgId,
		Data:     data,
		SyncPos:  constants.PSyncReqNeedSync,
		PushType: constants.PushTypeContent,
	}

	bytes, err := proto.Marshal(req)
	if err != nil {
		logger.Errorf("PChatMsgHandler syncToOther proto marshal %v failed:%v", req, err)
		return err
	}

	err = app.App.GetNatsClient().Publish(constants.MqSyncSubject, bytes)
	if err != nil {
		logger.Errorf("PChatMsgHandler syncToOther natsClient publish %v failed:%v", req, err)
		return err
	}
	return nil
}
