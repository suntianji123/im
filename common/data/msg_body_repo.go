package data

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/data/ent"
	"github.com/im/common/data/ent/msgbody"
	"github.com/im/common/util"
)

var MsgBodyRepo = msgBodyRepo{}

type msgBodyRepo struct {
}

func (p *msgBodyRepo) GetMsgBody(ctx context.Context, req *api.MsgBodyGetReq) (*api.MsgBodyGetResp, error) {
	if req.MsgIds == nil || len(req.MsgIds) == 0 {
		return nil, nil
	}

	keys := util.Util.ConvertInt64ToStrs(req.MsgIds)
	listCmd := DataM.GetRedisClient().MGet(ctx, keys...)
	arr, err := listCmd.Result()
	if err != nil {
		logger.Errorf("MsgBodyRepo GetMsgBody failed:%v", err)
		return nil, err
	}

	if len(arr) != len(keys) {
		return nil, constants.ErrIllegalState
	}

	m := make(map[int64]string, len(arr))
	for i, v := range arr {
		str := v.(string)
		if len(str) == 0 {
			continue
		}

		m[req.MsgIds[i]] = v.(string)
	}

	if len(req.MsgIds) > len(m) {
		dbIds := make([]int64, 0)
		for _, v := range req.GetMsgIds() {
			if _, ok := m[v]; !ok {
				dbIds = append(dbIds, v)
			}
		}

		dbs := p.fromDB(ctx, dbIds)
		for _, v := range dbs {
			m[v.MsgID] = v.Body
		}
	}
	return &api.MsgBodyGetResp{
		Bodies: m,
	}, nil
}

func (p *msgBodyRepo) fromDB(ctx context.Context, msgIds []int64) []*ent.MsgBody {
	return DataM.GetDBClient().MsgBody.Query().Where(msgbody.MsgIDIn(msgIds...)).AllX(ctx)
}
