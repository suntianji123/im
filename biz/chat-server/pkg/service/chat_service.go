package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/util"
)

type ChatServiceServerImpl struct {
	api.UnimplementedChatServiceServer
	msgChan chan *packet
}

type packet struct {
	handler Handler
	ctx     context.Context
	body    proto.Message
}

func NewChatServiceServerImpl() *ChatServiceServerImpl {
	p := &ChatServiceServerImpl{
		msgChan: make(chan *packet),
	}

	for i := 0; i < constants.GrouteLenth; i += 1 {
		go p.dispatcher()
	}

	return p
}

func (p *ChatServiceServerImpl) dispatcher() {
	for msg := range p.msgChan {
		err := msg.handler.Handle(msg.ctx, msg.body)
		if err != nil {
			logger.Errorf("ChatServiceServerImpl dispatcher handle %v failed:%v", msg, err)
		}
	}
}

func (p *ChatServiceServerImpl) Handle(ctx context.Context, req *api.HandleReq) (*api.Result, error) {
	uri, err := util.ProtocalUtil.GetUri(req.Data)
	if err != nil {
		logger.Errorf("ChatServiceServerImpl Handle protocalUtil GetUri data:%s faild:%v", req.Data, err)
		return nil, err
	}

	h, ok := Handlers[uri]
	if !ok {
		logger.Warnf("unknow uri:%d to handle", uri)
		return nil, nil
	}

	msg := h.Message()
	err = jsonpb.UnmarshalString(req.Data, msg)
	if err != nil {
		logger.Errorf("ChatServiceServerImpl Handle jsonpb unmarshalString data:%s failed:%v", req.Data, err)
		return nil, err
	}

	p.msgChan <- &packet{ctx: context.Background(), handler: h, body: msg}
	return &api.Result{}, nil
}
