package mq

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/mq"
	"github.com/im/common/session"
	"github.com/im/common/util"
	"github.com/im/signal/sync/pkg/service"
	"github.com/im/signal/sync/rpc"
	"github.com/nats-io/nats.go"
)

type syncSubHandler struct {
	mq.BaseSubHandler
}

type addr struct {
	ip   string
	port int
}

func (p *addr) address() string {
	return fmt.Sprintf("%s:%d", p.ip, p.port)
}

func (p *addr) key() string {
	return fmt.Sprintf("%s:%d", p.ip, p.port)
}

func NewSyncSubHandler() *syncSubHandler {
	return &syncSubHandler{
		mq.NewBaseSubHandler(constants.MqSyncSubject),
	}
}

func (p *syncSubHandler) Handle(msg *nats.Msg) error {
	req := &api.PSyncReq{}
	err := proto.Unmarshal(msg.Data, req)
	logger.Warnf("%s", string(msg.Data))
	if err != nil {
		logger.Errorf("syncSubHandler Handle %v failed:%v", msg, err)
		return err
	}

	ctx := context.Background()

	que := &service.Queue{Uid: req.Uid, Channel: req.Channel}
	b, err := service.QueueService.HasEnqueue(ctx, que, req)
	if err != nil {
		logger.Errorf("SyncHandler Handle QueueService hasEnqueue que:%s failed:%v", que.SyncKey(), err)
		return err
	}

	if !b {
		err = service.SyncPosService.SetSyncPos(ctx, que, req)
		if err != nil {
			logger.Errorf("SyncSubHandler Handle syncService setSyncPos queue:%s failed:%v", que.SyncKey(), err)
			return err
		}

		err = service.QueueService.Enqueue(ctx, que, req)
		if err != nil {
			logger.Errorf("SyncHandler Handle QueueService enqueue que:%s failed:%v", que.SyncKey(), err)
			return err
		}
	}

	//获取在线信息
	reqAndOnlines, err := service.OnlineService.GetOnline(ctx, req)
	if err != nil {
		logger.Errorf("SyncHandler Handle OnlineService getOnline %v failed:%v", req, err)
		return err
	}

	if len(reqAndOnlines.OnlineInfos) > 0 {
		//同步在线用户
		batches, err := p.makeBatch(p.makeDown(reqAndOnlines))
		if err != nil {
			logger.Errorf("SyncSubHandler Handle makeBatch failed:%v", err)
			return err
		}

		for ad, packets := range batches {
			for _, packet := range packets {
				err = rpc.BatchServiceClientImpl.Batch(ad.address(), packet)
				if err != nil {
					logger.Errorf("syncSubHandler batchServiceClientImpl batch failed:%v", err)
				}
			}
		}

	} else {
		//同步离线用户

	}
	return nil
}

func (p *syncSubHandler) makeBatch(addrAndDowns map[*addr][]*api.PTransDown) (map[*addr][]*api.PBatch, error) {
	mp := make(map[string]bool)
	result := make(map[*addr][]*api.PBatch)
	for add, list := range addrAndDowns {
		batches, err := util.ProtocalUtil.Pack(list)
		if err != nil {
			logger.Errorf("SyncSubHandler makeBatch failed:%v", err)
			return nil, err
		}

		key := add.key()
		if _, ok := mp[key]; ok {
			result[add] = append(result[add], batches...)
		} else {
			result[add] = append([]*api.PBatch{}, batches...)
			mp[key] = true
		}
	}
	return result, nil
}

func (p *syncSubHandler) makeDown(req *service.ReqAndOnlineInfos) map[*addr][]*api.PTransDown {
	result := make(map[*addr][]*api.PTransDown)
	mp := make(map[string]bool)
	for _, info := range req.OnlineInfos {
		a := &addr{ip: info.LoginIp, port: info.LoginPort}
		if _, ok := mp[a.key()]; ok {
			result[a] = append(result[a], p.makeTransDown(req.Req, info))
		} else {
			result[a] = []*api.PTransDown{p.makeTransDown(req.Req, info)}
			mp[a.key()] = true
		}
	}
	return result
}

func (p *syncSubHandler) makeTransDown(req *api.PSyncReq, info *session.OnlineInfo) *api.PTransDown {
	return &api.PTransDown{
		Uri:      constants.TransDown,
		AppId:    req.AppId,
		Uid:      req.Uid,
		DeviceId: info.DeviceId,
		Channel:  req.Channel,
		SyncPos:  req.SyncPos,
		MsgId:    req.MsgId,
		Data:     req.Data,
	}
}
