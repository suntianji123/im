package util

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/im/common/api"
	"github.com/im/common/constants"
)

var (
	ProtocalUtil   = &protocalUtil{}
	SidChatFormat  = "c:%d:%d:%d"
	SidGChatFormat = "g:%d:%d"
)

const maxLen = 64 * 1024

type protocalUtil struct {
}

func (p *protocalUtil) GetUri(data string) (int, error) {
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		logger.Errorf("ProtocalUtil deserialize json unmarshal data:%s failed:%v", data, err)
		return 0, err
	}

	if v, ok := m[constants.Uri]; ok {
		return int(v.(float64)), nil
	}
	return 0, nil
}

func (p *protocalUtil) GetChatSid(req *api.PChatMsgSendReq) string {
	a := req.FromUid
	b := req.ToUid
	if a > b {
		c := a
		a = b
		b = c
	}

	return fmt.Sprintf(SidChatFormat, a, b, req.Channel)
}

func (p *protocalUtil) Serialize(msg proto.Message) (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		logger.Errorf("ProtocalUtil Serialize json marshal data:%v failed:%v", msg, err)
		return "", err
	}
	return string(bytes), nil
}

func (p *protocalUtil) Pack(messages []*api.PTransDown) ([]*api.PBatch, error) {
	batchSize := 0
	batch := make([]string, 0)
	result := make([]*api.PBatch, 0)
	for _, msg := range messages {
		bytes, err := json.Marshal(msg)
		if err != nil {
			logger.Errorf("ProtocalUtil Pack %v failed:%v", msg, err)
			return nil, err
		}
		batchSize += len(bytes)
		if batchSize <= maxLen {
			batch = append(batch, string(bytes))
		} else {
			result = append(result, &api.PBatch{
				Data: append([]string{}, batch...),
			})
			clear(batch)
			batch = append(batch, string(bytes))
			batchSize = len(bytes)
		}
	}

	if len(batch) > 0 {
		result = append(result, &api.PBatch{
			Data: append([]string{}, batch...),
		})
	}
	return result, nil
}

func (p *protocalUtil) UnPack(batch *api.PBatch) ([]*api.PTransDown, error) {
	if batch.Data == nil || len(batch.Data) <= 0 {
		return nil, nil
	}

	result := make([]*api.PTransDown, len(batch.Data))
	for i, str := range batch.Data {
		down := &api.PTransDown{}
		err := json.Unmarshal([]byte(str), down)
		if err != nil {
			logger.Errorf("ProtocalUtil Unpack json umarshal %s failed:%v", str, err)
			return nil, err
		}
		result[i] = down
	}
	return result, nil
}

func (p *protocalUtil) GetPChatMsgSendReqJson(req *api.PChatMsgSendReq) *api.PChatMsgSendReqJson {
	return &api.PChatMsgSendReqJson{
		Uri:        req.Uri,
		AppId:      req.AppId,
		Channel:    req.Channel,
		MsgUuid:    req.MsgUuid,
		FromUid:    fmt.Sprintf("%d", req.FromUid),
		FromName:   req.FromName,
		ToUid:      fmt.Sprintf("%d", req.ToUid),
		ToAppId:    req.ToAppId,
		MsgType:    req.MsgType,
		Message:    req.Message,
		Cts:        req.Cts,
		DeviceType: req.DeviceType,
		DeviceId:   req.DeviceId,
		Extension:  req.Extension,
		MsgId:      fmt.Sprintf("%d", req.MsgId),
		SeqId:      req.SeqId,
		Sts:        req.Sts,
	}
}

func (p *protocalUtil) GetPTransDownJson(req *api.PTransDown) *api.PTransDownJson {
	return &api.PTransDownJson{
		Uri:      req.Uri,
		AppId:    req.AppId,
		Uid:      req.Uid,
		DeviceId: req.DeviceId,
		Channel:  req.Channel,
		SyncPos:  req.SyncPos,
		MsgId:    fmt.Sprintf("%d", req.MsgId),
		Data:     req.Data,
	}
}
