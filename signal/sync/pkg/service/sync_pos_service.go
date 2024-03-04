package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/data"
	"time"
)

var SyncPosService = &syncPosService{}

type syncPosService struct{}

func (p *syncPosService) SetSyncPos(ctx context.Context, que *Queue, req *api.PSyncReq) error {
	pipe := data.DataM.GetRedisClient().Pipeline()
	syncKey := que.SyncKey()
	current := time.Now().UnixMilli()
	err := pipe.SetNX(ctx, syncKey, fmt.Sprintf("%d", current), 0).Err()
	if err != nil {
		logger.Errorf("SyncPosService SetSyncPos pipeline setNx key:%s failed:%v", syncKey, err)
		return err
	}

	cmd := pipe.Incr(ctx, syncKey)
	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("SyncPosService SetSyncPos queue:%s failed:%v", syncKey, err)
		return err
	}

	if cmd.Err() != nil {
		logger.Errorf("SyncPosService SetSyncPos pipeline Incr key:%s failed:%v", syncKey, cmd.Err())
		return cmd.Err()
	}

	if req.SyncPos == constants.PSyncReqNeedSync {
		req.SyncPos = cmd.Val()
	}
	return nil
}
