package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/im/common/api"
	"github.com/im/common/data"
	"time"
)

const (
	ExpireSeconds = 7 * 24 * 60 * 60
)

var MsgBodyService = &msgBodyService{}

type msgBodyService struct {
}

func (p *msgBodyService) SaveBody(ctx context.Context, req *api.MsgBodySaveReq) error {
	err := data.DataM.GetRedisClient().SetEX(ctx, fmt.Sprintf("%d", req.MsgId), req.MsgBody, time.Duration(ExpireSeconds)*time.Second).Err()
	if err != nil {
		logger.Errorf("MsgBodyService SaveBody failed:%v", err)
		return err
	}
	return nil
}
