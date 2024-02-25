package data

import (
	"context"
	"github.com/im/common/data/ent"
)

var ImMsgRepo = &imMsgRepo{}

type imMsgRepo struct {
}

func (p *imMsgRepo) Create(ctx context.Context, imMsg *ent.IMMsg) *ent.IMMsg {
	return DataM.GetDBClient().IMMsg.Create().
		SetSid(imMsg.Sid).
		SetFromUID(imMsg.FromUID).
		SetFromAppid(imMsg.FromAppid).
		SetToUID(imMsg.ToUID).
		SetToAppid(imMsg.ToAppid).
		SetChannel(imMsg.Channel).
		SetMsgID(imMsg.MsgID).
		SetCts(imMsg.Cts).
		SaveX(ctx)
}
