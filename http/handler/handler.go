package handler

import "github.com/gin-gonic/gin"

var Handlers map[string]gin.HandlerFunc

func init() {
	Handlers = make(map[string]gin.HandlerFunc)
	friend := &friendHandler{}
	Handlers["/im/friend/list"] = friend.list

	im := &imHandler{}
	Handlers["/im/msg/sync"] = im.sync

	user := &userHandler{}
	Handlers["/im/user/login"] = user.login
	Handlers["/im/user/register"] = user.register
}
