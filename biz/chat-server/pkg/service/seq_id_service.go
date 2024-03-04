package service

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/im/biz/chat-server/app"
	"github.com/im/common/api"
	"github.com/im/common/data"
	"github.com/im/common/util"
	"golang.org/x/net/context"
	"time"
)

var SeqIdService = seqIdService{
	msgIdCache: util.NewExpiredMap(),
}

type seqIdService struct {
	msgIdCache *util.ExpiredMap
}

func (p *seqIdService) GetSeqId(ctx context.Context, req *api.PChatMsgSendReq) (int64, error) {
	sid := util.ProtocalUtil.GetChatSid(req)
	pipe := data.DataM.GetRedisClient().Pipeline()
	err := pipe.SetNX(ctx, sid, fmt.Sprintf("%d", time.Now().UnixMilli()), 0).Err()
	if err != nil {
		logger.Errorf("SeqIdService GetSeqId pipe setNX failed:%v", err)
		return 0, err
	}
	cmd := pipe.Incr(ctx, sid)

	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("SeqIdService GetSeqId pipe exec failed:%v", err)
		return 0, err
	}

	if cmd.Err() != nil {
		logger.Errorf("SeqIdService GetSeqId pipe incr failed:%v", cmd.Err())
		return 0, cmd.Err()
	}

	return cmd.Val(), nil
}

func (p *seqIdService) GetMsgId(req *api.PChatMsgSendReq) (int64, error) {
	if ok, v := p.msgIdCache.Get(req.MsgUuid); ok {
		return v.(int64), nil
	}

	msgId, err := app.App.GetIdGenerator().NextId()
	if err != nil {
		logger.Errorf("SeqIdService GetMsgId failed:%v", err)
		return 0, err
	}
	p.msgIdCache.Set(req.MsgUuid, msgId, 120)
	return msgId, nil
}
