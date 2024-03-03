package service

import (
	"context"
	"github.com/im/common/api"
)

type HeartBeatServiceServerImpl struct {
	api.UnimplementedHeartBeartServiceServer
}

func NewHeartBeatServiceServerImpl() *HeartBeatServiceServerImpl {
	return &HeartBeatServiceServerImpl{}
}

func (*HeartBeatServiceServerImpl) HeartBeart(context.Context, *api.PPing) (*api.PPong, error) {
	return &api.PPong{}, nil
}
