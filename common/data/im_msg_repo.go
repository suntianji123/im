package data

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/common/data/ent"
)

var ImMsgRepo = &imMsgRepo{}

type imMsgRepo struct {
}

func (p *imMsgRepo) Create(ctx context.Context, imMsg *ent.IMMsg) *ent.IMMsg {
	res, err := DataM.GetDBClient().IMMsg.Create().
		SetSid(imMsg.Sid).
		SetFromUID(imMsg.FromUID).
		SetFromAppid(imMsg.FromAppid).
		SetToUID(imMsg.ToUID).
		SetToAppid(imMsg.ToAppid).
		SetChannel(imMsg.Channel).
		SetMsgID(imMsg.MsgID).
		SetCts(imMsg.Cts).
		Save(ctx)
	if err != nil {
		logger.Errorf("ImMsgRepo Create failed:%v", err)
	}
	return res
}
