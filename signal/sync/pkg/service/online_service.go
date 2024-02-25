package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/im/common/api"
	"github.com/im/common/data"
	"github.com/im/common/session"
	"github.com/im/common/util"
	"math"
)

const emptyDeviceId = ""

var OnlineService = &onlineService{}

type onlineService struct{}

type UidAndAppId struct {
	Uid   int64
	AppId int
}

type PushDevice struct {
	Uaa      *UidAndAppId
	DeviceId string
}

type ReqAndPushDevices struct {
	Req     *api.PSyncReq
	Devices []*PushDevice
}

type ReqAndOnlineInfos struct {
	Req         *api.PSyncReq
	OnlineInfos []*session.OnlineInfo
}

func (p *onlineService) GetOnline(ctx context.Context, req *api.PSyncReq) (*ReqAndOnlineInfos, error) {
	reqAndPushDevices := &ReqAndPushDevices{
		Req:     req,
		Devices: p.CreatePushDevice(req),
	}

	//获取UidAppId对应的deviceId列表
	mp, err := p.searchOnlineResult(ctx, p.getUidAndAppIdKeyDeviceIds(reqAndPushDevices.Devices))
	if err != nil {
		logger.Errorf("onlineServcie GetOnline searchOnlineResult failed:%v", err)
		return nil, err
	}

	list := make([]*session.OnlineInfo, 0)
	for _, m1 := range mp {
		for _, info := range m1 {
			list = append(list, info)
		}
	}

	return &ReqAndOnlineInfos{Req: req, OnlineInfos: list}, nil
}

func (p *onlineService) searchOnlineResult(ctx context.Context, uAAToDeviceIdsMap map[*UidAndAppId][]string) (map[*UidAndAppId]map[string]*session.OnlineInfo, error) {
	pushAll := p.getPushAll(uAAToDeviceIdsMap)
	pushDevices := p.getPushDevices(uAAToDeviceIdsMap)
	pipe := data.DataM.GetRedisClient().Pipeline()

	mp := make(map[string]bool)
	otherDeviceIds := make([]string, 0)
	for _, uaa := range pushAll {
		list, err := pipe.ZRangeByScoreWithScores(ctx, p.userKey(uaa.Uid, uaa.AppId), &redis.ZRangeBy{
			Min: fmt.Sprintf("%f", math.Inf(-1)),
			Max: fmt.Sprintf("%f", math.Inf(1)),
		}).Result()
		if err != nil {
			logger.Errorf("OnlineService searchOnlineResult pipeline ZRangeByScoreWithScores key:%s failed:%v", p.userKey(uaa.Uid, uaa.AppId), err)
			return nil, err
		}

		for _, r := range list {
			key := p.deviceKey(uaa.Uid, uaa.AppId, r.Member.(string))
			if _, ok := mp[key]; !ok {
				otherDeviceIds = append(otherDeviceIds, key)
				mp[key] = true
			}
		}
	}

	clear(mp)
	for _, d := range otherDeviceIds {
		if _, ok := mp[d]; !ok {
			pushDevices = append(pushDevices, d)
			mp[d] = true
		}
	}

	list, err := pipe.MGet(ctx, pushDevices...).Result()
	if err != nil {
		logger.Errorf("OnlineService searchOnlineResult pipe mget %v failed:%v", pushDevices, err)
		return nil, err
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("OnlineService searchOnlineResult pipe exec failed:%v", err)
		return nil, err
	}

	clear(mp)
	result := make(map[*UidAndAppId]map[string]*session.OnlineInfo)
	for _, str := range list {
		info := &session.OnlineInfo{}
		err = json.Unmarshal([]byte(str.(string)), info)
		if err != nil {
			logger.Errorf("OnlineService searchOnlineResult json Unmarshal %v failed:%v", str, err)
			return nil, err
		}

		userKey := p.userKey(info.Uid, info.AppId)
		if _, ok := mp[userKey]; ok {
			result[&UidAndAppId{Uid: info.Uid, AppId: info.AppId}][info.DeviceId] = info
		} else {
			result[&UidAndAppId{Uid: info.Uid, AppId: info.AppId}] = map[string]*session.OnlineInfo{info.DeviceId: info}
			mp[userKey] = true
		}
	}
	return result, nil
}

func (p *onlineService) getPushAll(uAAToDeviceIdsMap map[*UidAndAppId][]string) []*UidAndAppId {
	result := make([]*UidAndAppId, 0)
	for uaa, deviceIds := range uAAToDeviceIdsMap {
		exists := false
		for _, deviceId := range deviceIds {
			if deviceId == emptyDeviceId {
				exists = true
				break
			}
		}

		if exists {
			result = append(result, uaa)
		}
	}
	return result
}

func (p *onlineService) getPushDevices(uAAToDeviceIdsMap map[*UidAndAppId][]string) []string {
	result := make([]string, 0)
	mp := make(map[string]bool)
	for uaa, deviceIds := range uAAToDeviceIdsMap {
		for _, deviceId := range deviceIds {
			if _, ok := mp[deviceId]; !ok && deviceId != emptyDeviceId {
				result = append(result, p.deviceKey(uaa.Uid, uaa.AppId, deviceId))
				mp[deviceId] = true
			}
		}

	}
	return result
}

func (p *onlineService) userKey(uid int64, appId int) string {
	return fmt.Sprintf("ol:{%d}:%d", uid, appId)
}

func (p *onlineService) deviceKey(uid int64, appId int, deviceId string) string {
	return fmt.Sprintf("ol:{%d}:%d:%s", uid, appId, deviceId)
}

func (p *onlineService) getUidAndAppIdKeyDeviceIds(pushDevices []*PushDevice) map[*UidAndAppId][]string {
	result := make(map[*UidAndAppId][]string)
	mp := make(map[string]bool)
	for _, pushDevice := range pushDevices {
		if _, ok := mp[p.userKey(pushDevice.Uaa.Uid, pushDevice.Uaa.AppId)]; ok {
			result[pushDevice.Uaa] = append(result[pushDevice.Uaa], pushDevice.DeviceId)
		} else {
			result[pushDevice.Uaa] = []string{pushDevice.DeviceId}
			mp[p.userKey(pushDevice.Uaa.Uid, pushDevice.Uaa.AppId)] = true
		}
	}
	return result
}

func (p *onlineService) CreatePushDevice(req *api.PSyncReq) []*PushDevice {
	uaa := &UidAndAppId{
		Uid:   req.Uid,
		AppId: int(req.AppId),
	}
	if len(req.DeviceId) > 0 {
		return []*PushDevice{
			{
				Uaa:      uaa,
				DeviceId: req.DeviceId,
			},
		}
	}

	appIds := util.ChannelUtil.GetAppIds(int(req.ChatType), int(req.Channel), int(req.AppId))
	result := make([]*PushDevice, len(appIds))
	for i, appId := range appIds {
		result[i] = &PushDevice{
			Uaa:      &UidAndAppId{Uid: req.Uid, AppId: appId},
			DeviceId: emptyDeviceId,
		}
	}
	return result
}
