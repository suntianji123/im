package data

import (
	"context"
	"github.com/im/common/data/ent"
	"github.com/im/common/data/ent/msgbody"
)

var QueueRepo = &queueRepo{}

type queueRepo struct {
}

func (p *queueRepo) GetMsgBody(ids []int64) []*ent.MsgBody {
	return DataM.GetDBClient().MsgBody.Query().Where(msgbody.IDIn(ids...)).AllX(context.Background())
}
