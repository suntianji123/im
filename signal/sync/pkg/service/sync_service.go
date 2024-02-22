package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/im/common/api"
)

type SyncServiceServerImpl struct {
	api.UnimplementedSyncServiceServer
}

func init() {
	config.SetProviderService(&SyncServiceServerImpl{})
}

func (*SyncServiceServerImpl) Sync(ctx context.Context, req *api.SyncReq) (*api.Result, error) {
	result, err := QueueService.SyncMsg(ctx, &Queue{uid: req.Uid,
		channel: req.Channel,
	}, req)

	if err != nil {
		logger.Errorf("SyncServiceServerImpl Sync failed:%v", err)
		return nil, err
	}
	return result, nil
}
