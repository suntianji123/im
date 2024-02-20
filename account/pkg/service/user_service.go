package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/im/account/rpc"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/data"
	"google.golang.org/protobuf/types/known/anypb"
)

type UserServiceServerImpl struct {
	api.UnimplementedUserServiceServer
}

func (*UserServiceServerImpl) Login(ctx context.Context, req *api.UserLoginReq) (*api.Result, error) {
	user := data.UserRepo.FindByUsernameAndPassword(ctx, req.Username, req.Password)
	if user == nil {
		return &api.Result{
			Code: constants.ErrorCode,
			Msg:  "用户不存在",
		}, nil
	}

	remote, err := rpc.AuthServiceClientImpl.Auth(ctx)
	if err != nil {
		logger.Errorf("远程请求auth失败:%v", err)
		return nil, err
	}

	err = remote.Send(&api.AuthReq{
		Uid:        user.ID,
		AppId:      req.AppId,
		DeviceType: req.DeviceType,
	})
	if err != nil {
		logger.Errorf("向auth服务鉴权失败:%v", err)
		return nil, err
	}

	authResult := &api.Result{}
	err = remote.RecvMsg(authResult)
	if err != nil {
		logger.Errorf("接收auth服务鉴权失败:%v", err)
		return nil, err
	}

	authResp, err := authResult.Data.UnmarshalNew()
	if err != nil {
		logger.Errorf("反序列化AuthResp失败:%v", err)
		return nil, err
	}

	data, err := anypb.New(&api.UserLoginResp{
		UserInfo: &api.UserInfo{
			Id:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Nickname,
			Ext:      user.Ext,
		},
		Token: authResp.(*api.AuthResp).Token,
	})
	if err != nil {
		logger.Errorf("%v", err)
		return nil, err
	}
	return &api.Result{Code: 0, Data: data}, nil
}

func (*UserServiceServerImpl) Register(ctx context.Context, req *api.UserRegisterReq) (*api.Result, error) {
	if data.UserRepo.CountByUsername(ctx, req.Username) > 0 {
		return &api.Result{
			Code: constants.ErrorCode,
			Msg:  "用户已存在",
		}, nil
	}

	data, err := anypb.New(&api.UserRegisterResp{
		Uid: data.UserRepo.Create(ctx, req).ID,
	})
	if err != nil {
		logger.Errorf("UserServiceServerImpl Register failed:%v", err)
		return nil, err
	}

	return &api.Result{Data: data}, nil
}

func init() {
	config.SetProviderService(&UserServiceServerImpl{})
}
