package service

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/dubbogo/grpc-go/codes"
	"github.com/dubbogo/grpc-go/status"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/util"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

type AuthServiceServerImpl struct {
	api.UnimplementedAuthServiceServer
}

func (p *AuthServiceServerImpl) Auth(stream api.AuthService_AuthServer) error {
	req, err := stream.Recv()
	if err != nil {
		logger.Errorf("Auth流接收消息失败:%v", err)
		return err
	}

	token, err := p.createToken(req)
	if err != nil {
		logger.Errorf("AuthServiceServerImpl Auth failed:%v", err)
		return err
	}

	any, err := anypb.New(&api.AuthResp{
		Token: token,
	})
	if err != nil {
		logger.Errorf("Auth流序列号any消息失败:%v", err)
		return err
	}

	stream.SendMsg(&api.Result{Code: 0, Data: any})

	return status.Errorf(codes.Unimplemented, "method Auth not implemented")
}

func (*AuthServiceServerImpl) CheckAuth(stream api.AuthService_CheckAuthServer) error {
	req, err := stream.Recv()
	if err != nil {
		logger.Errorf("AuthServiceServerImpl CheckAuth failed:%v", err)
		return err
	}

	payload, err := util.JwtUtil.Parse("", req.Token)
	if err != nil {
		logger.Errorf("AuthServiceServerImpl CheckAuth failed:%v", err)
		return err
	}

	if payload.Uid == req.Uid && payload.AppId == int(req.AppId) && payload.DeviceType == int(req.DeviceType) {
		err = stream.SendMsg(&api.Result{Code: constants.SuccessCode})
	} else {
		err = stream.SendMsg(&api.Result{Code: constants.ErrorCode, Msg: "token illegal"})
	}

	if err != nil {
		logger.Errorf("AuthServiceServerImpl CheckAuth failed:%v", err)
		return err
	}
	return nil
}

func (*AuthServiceServerImpl) createToken(req *api.AuthReq) (string, error) {
	token, err := util.JwtUtil.Encode("", util.NewPayload(req.Uid, int(req.AppId), int(req.DeviceType)))
	if err != nil {
		logger.Errorf("AuthServiceServerImpl createToken failed:%v", err)
		return "", err
	}

	return token, nil
}

func init() {
	config.SetProviderService(&AuthServiceServerImpl{})
}
