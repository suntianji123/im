package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/im/common/api"
	"github.com/im/common/data"
	"google.golang.org/protobuf/types/known/anypb"
)

type MessageServiceServerImpl struct {
	api.UnimplementedMessageServiceServer
}

func init() {
	config.SetProviderService(&MessageServiceServerImpl{})
}

func (*MessageServiceServerImpl) GetMsgBody(ctx context.Context, req *api.MsgBodyGetReq) (*api.Result, error) {

	resp, err := data.MsgBodyRepo.GetMsgBody(ctx, req)
	if err != nil {
		logger.Errorf("MessageServiceServerImpl GetMsgBody failed:%v", err)
		return nil, err
	}

	any, err := anypb.New(&api.MsgDataGetResp{
		Msgs: resp.Bodies,
	})
	if err != nil {
		logger.Errorf("MessageServiceServerImpl GetMsgBody anyPb failed:%v", err)
		return nil, err
	}

	return &api.Result{Data: any}, nil
}

func (*MessageServiceServerImpl) SaveMsgBody(ctx context.Context, req *api.MsgBodySaveReq) (*api.Result, error) {
	err := MsgBodyService.SaveBody(ctx, req)
	if err != nil {
		logger.Errorf("MessageService SaveMsgBody:%s failed:%v", req.MsgBody, err)
		return nil, err
	}
	return &api.Result{}, nil
}
