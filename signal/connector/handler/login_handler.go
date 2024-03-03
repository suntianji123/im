package handler

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	s "github.com/im/common/session"
	"github.com/im/common/util"
	"github.com/im/signal/connector/app"
	"github.com/im/signal/connector/pkg/service"
	"github.com/im/signal/connector/rpc"
	"time"
)

type loginHandler struct {
}

func (p *loginHandler) Message() proto.Message {
	return &api.PLoginReq{}
}

func (p *loginHandler) Handler(ctx context.Context, session *s.Session, msg proto.Message) error {
	req := msg.(*api.PLoginReq)
	onlineInfo := session.GetOnlineInfo()

	if onlineInfo == nil {
		err := p.doLogin(ctx, session, req)
		if err != nil {
			logger.Errorf("LoginHandler handler doLogin failed:%v", err)
			return err
		}
	} else {
		err := p.reLogin(ctx, session, req, onlineInfo)
		if err != nil {
			logger.Errorf("LoginHandler Handler relogin failed:%v", err)
			return err
		}
	}
	return nil
}

func (p *loginHandler) doLogin(ctx context.Context, session *s.Session, msg *api.PLoginReq) error {
	stream, err := rpc.AuthServiceClientImpl.CheckAuth(ctx)
	if err != nil {
		logger.Errorf("LoginHandler Handler CheckAuth get stream failed:%v", err)
		return err
	}

	err = stream.Send(&api.AuthCheckReq{
		Uid:        msg.Uid,
		AppId:      msg.AppId,
		DeviceType: msg.DeviceType,
		Token:      msg.Token,
	})
	if err != nil {
		logger.Errorf("LoginHandler doLogin stream send failed:%v", err)
		return err
	}

	var result = &api.Result{}
	err = stream.RecvMsg(result)
	if err != nil {
		logger.Errorf("LoginHandler doLogin stream receive failed:%v", err)
		return err
	}

	if result.Code == constants.SuccessCode {
		err = p.onLoginOk(ctx, session, msg)
		if err != nil {
			logger.Errorf("LoginHandler doLogin onLoginOk failed:%v", err)
			return err
		}
	}

	err = session.Push(ctx, &api.PLoginResp{
		Uri:       constants.LoginResp,
		Uid:       msg.Uid,
		AppId:     msg.AppId,
		RequestId: msg.RequestId,
		Code:      result.Code,
		Cause:     result.Msg,
	})

	if err != nil {
		logger.Errorf("LoginHandler doLogin Push failed:%v", err)
		return err
	}
	return nil
}

func (p *loginHandler) genOnlineInfo(session *s.Session, req *api.PLoginReq) (*s.OnlineInfo, error) {
	return &s.OnlineInfo{
		Uid:       req.Uid,
		AppId:     int(req.AppId),
		DeviceId:  req.DeviceId,
		LoginIp:   util.NetUtil.Ipv4(),
		LoginPort: app.App.GetBuilder().GetGrpcServerConfig().Port,
		LoginTs:   time.Now().UnixMilli(),
		LoginUuid: uuid.New().String(),
		SocketId:  fmt.Sprintf("%d", session.ID()),
	}, nil
}

func (p *loginHandler) onLoginOk(ctx context.Context, session *s.Session, req *api.PLoginReq) error {
	onlineInfo, err := p.genOnlineInfo(session, req)
	if err != nil {
		logger.Errorf("LoginHandler onLoginOk genOnlineInfo failed:%v", err)
		return err
	}
	err = service.OnlineService.AddOnline(ctx, onlineInfo)
	if err != nil {
		logger.Errorf("LoginHandler onLoginOk addOnline failed:%v", err)
		return err
	}

	session.GetAndSetOnlineInfo(onlineInfo)
	pre := session.BindSession(onlineInfo)
	if pre != nil && pre.ID() != session.ID() && pre.IsActive() {
		err = pre.Push(ctx, &api.PKickOff{
			Uri:  constants.KickOff,
			Code: constants.RespCodeReplaceByOther,
			Msg:  fmt.Sprintf("replace by other session:%s", session.RemoteAddr().String()),
		})
		if err != nil {
			logger.Errorf("LoginHander onloginOk push to pre session failed:%v", err)
			return nil
		}
		err = pre.Close()
		if err != nil {
			logger.Errorf("LoginHandler onLoginOk close pre session failed:%v", err)
			return nil
		}
	}
	return nil
}

func (p *loginHandler) reLogin(ctx context.Context, session *s.Session, req *api.PLoginReq, onlineInfo *s.OnlineInfo) error {
	if p.notSame(req, onlineInfo) {
		err := session.Push(ctx, p.forbiden(req))
		if err != nil {
			logger.Errorf("LoginHandler relogin session push failed:%v", err)
			return err
		}
	}

	logger.Infof("duplicate login req:%v", req)
	err := p.doLogin(ctx, session, req)
	if err != nil {
		logger.Errorf("LoginHandler reLogin dologin failed:%v", err)
		return err
	}
	return nil
}

func (p *loginHandler) notSame(req *api.PLoginReq, info *s.OnlineInfo) bool {
	return req.Uid != info.Uid || int(req.AppId) != info.AppId || req.DeviceId != info.DeviceId
}

func (p *loginHandler) forbiden(req *api.PLoginReq) *api.PLoginResp {
	return &api.PLoginResp{
		Uri:       constants.LoginResp,
		Uid:       req.Uid,
		AppId:     req.AppId,
		RequestId: req.RequestId,
		Code:      constants.IllegalParam,
		Cause:     "不能使用不同参数重复登陆",
	}
}
