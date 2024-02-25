package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/biz/chat-server/rpc"
	"github.com/im/common/api"
	"github.com/im/common/util"
)

var MsgService = &msgService{}

type msgService struct {
}

func (*msgService) SaveBody(ctx context.Context, req *api.PChatMsgSendReq) error {
	msgBody, err := util.ProtocalUtil.Serialize(req)
	if err != nil {
		logger.Errorf("MsgService SaveBody ProtocalUtil Serailize failed:%v", err)
		return err
	}

	_, err = rpc.MessageServiceClientImpl.SaveMsgBody(ctx, &api.MsgBodySaveReq{MsgId: req.MsgId,
		MsgBody: msgBody,
	})

	if err != nil {
		logger.Errorf("MsgService SaveBody rpc messsageServiceClientImpl saveMsgBody:%s failed:%v", msgBody, err)
		return err
	}

	return nil
}
