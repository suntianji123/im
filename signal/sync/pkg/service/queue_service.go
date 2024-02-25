package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/im/common/api"
	"github.com/im/common/data"
	"github.com/im/signal/sync/rpc"
	"google.golang.org/protobuf/types/known/anypb"
	"math"
	"strconv"
	"strings"
)

var QueueService = &queueService{}

type queueService struct {
}

var QueueRepo = &queueRepo{}

type Queue struct {
	Uid     int64
	Channel int32
}

type PeekResult struct {
	Last     *api.SyncMember
	Elements []*api.SyncMember
}

func (p *Queue) SyncKey() string {
	return fmt.Sprintf("sync:%d:%d", p.Uid, p.Channel)
}

func (p *Queue) queueKey() string {
	return fmt.Sprintf("queue:%d:%d", p.Uid, p.Channel)
}

type queueRepo struct {
}

func (p *queueService) SyncMsg(ctx context.Context, queue *Queue, syncReq *api.SyncReq) (*api.Result, error) {
	peekResult, err := p.PeekQueue(ctx, queue, syncReq.LocalSyncPos)
	if err != nil {
		logger.Errorf("QueueRepo SyncMsg failed:%v", err)
		return nil, err
	}
	if peekResult.Elements == nil || len(peekResult.Elements) == 0 {
		any, err := anypb.New(&api.SyncResp{
			Members: make([]*api.SyncMember, 0),
		})
		if err != nil {
			logger.Errorf("QueueRepo SyncMsg failed:%v", err)
			return nil, err
		}
		return &api.Result{Data: any}, nil
	}

	result, err := rpc.MessageServiceClientImpl.GetMsgBody(ctx, &api.MsgBodyGetReq{MsgIds: p.getMsgIds(peekResult.Elements)})
	if err != nil {
		logger.Errorf("QueueRepo SyncMsg GetMsgBody failed:%v", err)
		return nil, err
	}

	data := &api.MsgDataGetResp{}
	err = result.Data.UnmarshalTo(data)
	if err != nil {
		logger.Errorf("QueueRepo SyncMsg UnmarshalNew failed:%v", err)
		return nil, err
	}

	resp := &api.SyncResp{
		Members:  peekResult.Elements,
		Messages: data.Msgs,
		HasMore:  peekResult.Last != nil && len(peekResult.Elements) > 0 && peekResult.Last.GetSyncPos() > peekResult.Elements[len(peekResult.Elements)-1].GetSyncPos(),
	}

	any, err := anypb.New(resp)
	if err != nil {
		logger.Errorf("QueueRepo SyncMsg anyPb new failed:%v", err)
		return nil, err
	}
	return &api.Result{
		Data: any,
	}, nil
}

func (p *queueService) PeekQueue(ctx context.Context, queue *Queue, localSyncPos int64) (*PeekResult, error) {
	pipe := data.DataM.GetRedisClient().Pipeline()
	key := queue.queueKey()
	rangeResp := pipe.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min: fmt.Sprintf("%d", localSyncPos+1),
		Max: fmt.Sprintf("%d", math.MaxInt64),
	})

	lastResp := pipe.ZRange(ctx, key, -1, -1)

	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("QueueRepo PeekQueue failed:%v", err)
		return nil, err
	}

	elements, err := rangeResp.Result()
	if err != nil {
		logger.Errorf("QueueRepo PeekQueue rangeResp failed:%v", err)
		return nil, err
	}

	last, err := lastResp.Result()
	if err != nil {
		logger.Errorf("QueueRepo PeekQueue lastResp failed:%v", err)
		return nil, err
	}

	memebers, err := p.fromStrings(elements)
	if err != nil {
		logger.Errorf("QueueRepo PeekQueue lastResp failed:%v", err)
		return nil, err
	}

	lastElements, err := p.fromStrings(last)
	if err != nil {
		logger.Errorf("QueueRepo PeekQueue lastResp failed:%v", err)
		return nil, err
	}
	var lastElement *api.SyncMember
	if lastElements == nil || len(lastElements) == 0 {
		lastElement = nil
	} else {
		lastElement = lastElements[0]
	}

	return &PeekResult{Last: lastElement, Elements: memebers}, nil

}

func (p *queueService) fromStrings(strs []string) ([]*api.SyncMember, error) {
	if strs == nil || len(strs) <= 0 {
		return nil, nil
	}

	members := make([]*api.SyncMember, len(strs))
	for i, str := range strs {
		arr := strings.Split(str, ":")
		pos, err := strconv.ParseInt(arr[0], 10, 64)
		if err != nil {
			logger.Errorf("QueueRepo fromStrings failed:%v", err)
			return nil, err
		}

		msgId, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			logger.Errorf("QueueRepo fromStrings failed:%v", err)
			return nil, err
		}

		members[i] = &api.SyncMember{
			SyncPos: pos,
			MsgId:   msgId,
		}
	}
	return members, nil
}

func (p *queueService) getMsgIds(members []*api.SyncMember) []int64 {
	msgIds := make([]int64, len(members))
	for i, v := range members {
		msgIds[i] = v.MsgId
	}
	return msgIds
}

func (p *queueService) Enqueue(ctx context.Context, que *Queue, req *api.PSyncReq) error {
	queueKey := que.queueKey()
	pipe := data.DataM.GetRedisClient().Pipeline()
	err := pipe.ZAdd(ctx, queueKey, &redis.Z{
		Member: p.MemberString(req),
		Score:  float64(req.SyncPos),
	}).Err()
	if err != nil {
		logger.Errorf("queueService Enqueue member:%s score:%d pipe zadd failed:%v", queueKey, req.SyncPos, err)
		return err
	}

	err = pipe.ZRemRangeByRank(ctx, queueKey, 0, -1001).Err()
	if err != nil {
		logger.Errorf("queueService Enqueue member:%s score:%d pipe ZRemRangeByRank failed:%v", queueKey, req.SyncPos, err)
		return err
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("queueService Enqueue member:%s score:%d pipe exec failed:%v", queueKey, req.SyncPos, err)
		return err
	}
	return nil
}

func (p *queueService) MemberString(req *api.PSyncReq) string {
	return fmt.Sprintf("%d:%d", req.SyncPos, req.MsgId)
}
