package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/data"
	"strconv"
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

	err = pipe.Incr(ctx, syncKey).Err()
	if err != nil {
		logger.Errorf("SyncPosService SetSyncPos pipelien Incr key:%s failed:%v", syncKey, err)
		return err
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("SyncPosService SetSyncPos queue:%s failed:%v", syncKey, err)
		return err
	}

	str, err := data.DataM.GetRedisClient().Get(ctx, syncKey).Result()
	if err != nil {
		logger.Errorf("SyncPosService SetSyncPos pipelien get key:%s failed:%v", syncKey, err)
		return err
	}

	if req.SyncPos == constants.PSyncReqNeedSync {
		v1, err1 := strconv.ParseInt(str, 10, 64)
		if err1 != nil {
			logger.Errorf("SyncPosService SetSyncPos strconv parse %s failed:%v", str, err)
			return err1
		}
		req.SyncPos = v1
	}
	return nil
}
