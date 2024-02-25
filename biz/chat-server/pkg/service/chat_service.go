package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/golang/protobuf/jsonpb"
	"github.com/im/common/api"
	"github.com/im/common/util"
)

type ChatServiceServerImpl struct {
	api.UnimplementedChatServiceServer
}

func (*ChatServiceServerImpl) Handle(ctx context.Context, req *api.HandleReq) (*api.Result, error) {
	uri, err := util.ProtocalUtil.GetUri(req.Data)
	if err != nil {
		logger.Errorf("ChatServiceServerImpl Handle protocalUtil GetUri data:%s faild:%v", req.Data, err)
		return nil, err
	}

	h, ok := Handlers[uri]
	if !ok {
		logger.Warnf("unknow uri:%d to handle", uri)
		return nil, nil
	}

	msg := h.Message()
	err = jsonpb.UnmarshalString(req.Data, msg)
	if err != nil {
		logger.Errorf("ChatServiceServerImpl Handle jsonpb unmarshalString data:%s failed:%v", req.Data, err)
		return nil, err
	}
	err = h.Handle(ctx, msg)
	if err != nil {
		logger.Errorf("ChatServiceServerImpl handle uri:%d failed:%v", uri, err)
		return nil, err
	}
	return &api.Result{}, nil
}
