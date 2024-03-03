package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/common/api"
	"github.com/im/common/util"
	"github.com/im/signal/connector/app"
)

type BatchServiceServerImpl struct {
	api.UnimplementedBatchServiceServer
}

func NewBatchServiceServerImpl() *BatchServiceServerImpl {
	return &BatchServiceServerImpl{}
}

func (*BatchServiceServerImpl) Batch(ctx context.Context, batch *api.PBatch) (*api.Result, error) {
	downs, err := util.ProtocalUtil.UnPack(batch)
	if err != nil {
		logger.Errorf("BatchServiceServerImpl Batch ProtocalUtil Unpack %v failed:%v", batch, err)
		return nil, err
	}

	if downs != nil && len(downs) > 0 {
		for _, down := range downs {
			s := app.App.GetTcpAcceptor().GetAgentFactor().GetSessionPool().GetSession(down.Uid, int(down.AppId), down.DeviceId)
			if s == nil {
				logger.Warnf("BatchServiceServerImpl Batch Session %d:%d:%s is closed,push message:%v failed", down.Uid, down.AppId, down.DeviceId, down)
				continue
			}
			err = s.Push(ctx, util.ProtocalUtil.GetPTransDownJson(down))
			if err != nil {
				logger.Errorf("BatchServiceServerImpl Batch session %d:%d:%s push message %v failed:%v", down.Uid, down.AppId, down.DeviceId, down, err)
			}
		}
	}
	return &api.Result{}, nil
}
