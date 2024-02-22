package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/im/common/session"
	"github.com/im/signal/connector/app"
	"time"
)

var OnlineService = onlineService{}

const (
	userKeyFormat   = "ol:{%d}:%d"
	deviceKeyFormat = "ol:{%d}:%d:%s"
	expire          = 30 * 24 * 60 * 60
)

type onlineService struct{}

func (p *onlineService) AddOnline(ctx context.Context, info *session.OnlineInfo) error {
	pipe := app.App.GetData().GetRedisClient().Pipeline()
	userKey := p.userKey(info)
	_, err := pipe.ZAdd(ctx, userKey, &redis.Z{Score: float64(info.LoginTs), Member: p.userValue(info)}).Result()
	if err != nil {
		logger.Errorf("OnlineService AddOnline failed:%v", err)
		return err
	}

	duration := time.Duration(expire) * time.Second
	_, err = pipe.Expire(ctx, userKey, duration).Result()
	if err != nil {
		logger.Errorf("OnlineService AddOnline pipe expired failed:%v", err)
		return err
	}

	deviceValue, err := p.deviceValue(info)
	if err != nil {
		logger.Errorf("OnlineService AddOnline deviceValue failed:%v", err)
		return err
	}

	_, err = pipe.SetEX(ctx, p.deviceKey(info), deviceValue, duration).Result()
	if err != nil {
		logger.Errorf("OnlineServcie AddOnline pipeline setex failed:%v", err)
		return err
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("OnlineService AddOnline Pipeline exec failed:%v", err)
		return err
	}

	logger.Infof("OnlineService AddOnline success:%v", info)
	return nil
}

func (p *onlineService) userKey(onlineInfo *session.OnlineInfo) string {
	return fmt.Sprintf(userKeyFormat, onlineInfo.Uid, onlineInfo.AppId)
}

func (p *onlineService) deviceKey(onlineInfo *session.OnlineInfo) string {
	return fmt.Sprintf(deviceKeyFormat, onlineInfo.Uid, onlineInfo.AppId, onlineInfo.DeviceId)
}

func (p *onlineService) userValue(onlineInfo *session.OnlineInfo) string {
	return onlineInfo.DeviceId
}

func (p *onlineService) deviceValue(onlineInfo *session.OnlineInfo) (string, error) {
	bytes, err := json.Marshal(onlineInfo)
	if err != nil {
		logger.Errorf("OnlineService deviceValue json marshal failed:%v", err)
		return "", err
	}
	return string(bytes), nil
}
