package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/im/common/api"
	"github.com/im/common/data"
	"google.golang.org/protobuf/types/known/anypb"
)

type FriendServiceServerImpl struct {
	api.UnimplementedFriendServiceServer
}

func init() {
	config.SetProviderService(&FriendServiceServerImpl{})
}

func (*FriendServiceServerImpl) FriendList(ctx context.Context, req *api.FriendListReq) (*api.Result, error) {
	users, err := data.FriendRepo.FindFriends(ctx, req)
	if err != nil {
		logger.Errorf("FriendServiceServerImpl FriendList failed:%v", err)
		return nil, err
	}

	var friends []*api.UserInfo
	if users != nil && len(users) > 0 {
		friends = make([]*api.UserInfo, len(users))
		for i, v := range users {
			friends[i] = data.FriendRepo.ConvertEntUserInfoToPbUserInfo(v)
		}
	}

	any, err := anypb.New(&api.FriendListResp{
		Friends: friends,
	})
	if err != nil {
		logger.Errorf("FriendServiceServerImpl FriendList failed:%v", err)
		return nil, err
	}

	return &api.Result{
		Data: any,
	}, nil
}
