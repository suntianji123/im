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
	err := pipe.SetNX(ctx, syncKey, fmt.Sprintf("%d", time.Now().UnixMilli()), 0).Err()
	if err != nil {
		logger.Errorf("SyncPosService SetSyncPos pipeline setNx key:%s failed:%v", syncKey, err)
		return err
	}
	v, err := pipe.Incr(ctx, syncKey).Result()
	if err != nil {
		logger.Errorf("SyncPosService SetSyncPos pipelien Incr key:%s failed:%v", syncKey, err)
		return err
	}

	if req.SyncPos == constants.PSyncReqNeedSync {
		req.SyncPos = v
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("SyncPosService SetSyncPos queue:%s failed:%v", syncKey, err)
		return err
	}
	return nil
}
