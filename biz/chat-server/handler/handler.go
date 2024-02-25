package handler

import (
	"github.com/im/biz/chat-server/pkg/service"
	"github.com/im/common/constants"
)

func init() {
	service.Handlers[constants.ChatMsg] = &pChatMsgHandler{}
}
