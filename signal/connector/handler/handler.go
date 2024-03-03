package handler

import (
	"github.com/im/common/acceptor"
	"github.com/im/common/constants"
)

var Handlers = make(map[int]acceptor.TcpHandler)

func init() {
	Handlers[constants.Ping] = &heartBeatHandler{}
	Handlers[constants.LoginReq] = &loginHandler{}
	Handlers[constants.TransUp] = &pChatMsgHandler{}
}
