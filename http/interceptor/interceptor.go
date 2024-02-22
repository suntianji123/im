package interceptor

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/gin-gonic/gin"
	"github.com/im/common/api"
	"github.com/im/http/rpc"
	"strconv"
)

var AuthIntecetor gin.HandlerFunc = func(c *gin.Context) {
	path := c.Request.RequestURI
	if path == "/im/user/login" || path == "/im/user/register" {
		c.Next()
		return
	}

	uidStr := c.GetHeader("uid")
	if len(uidStr) == 0 {
		c.JSON(401, "")
		return
	}

	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		logger.Errorf("Auth Interceptor failed:%v", err)
		return
	}

	appIdStr := c.GetHeader("appId")
	if len(uidStr) == 0 {
		c.JSON(401, "")
		return
	}

	appId, err := strconv.ParseInt(appIdStr, 10, 32)
	if err != nil {
		logger.Errorf("Auth Interceptor failed:%v", err)
		return
	}

	deviceTypeStr := c.GetHeader("deviceType")
	if len(uidStr) == 0 {
		c.JSON(401, "")
		return
	}

	deviceType, err := strconv.ParseInt(deviceTypeStr, 10, 32)
	if err != nil {
		logger.Errorf("Auth Interceptor failed:%v", err)
		return
	}

	token := c.GetHeader("token")
	if len(uidStr) == 0 {
		c.JSON(401, "")
		return
	}

	stream, err := rpc.AuthServiceClientImpl.CheckAuth(c)
	if err != nil {
		logger.Errorf("Auth Interceptor failed:%v", err)
		return
	}

	err = stream.Send(&api.AuthCheckReq{
		Uid:        uid,
		AppId:      int32(appId),
		DeviceType: int32(deviceType),
		Token:      token,
	})
	if err != nil {
		logger.Errorf("Auth Interceptor failed:%v", err)
		return
	}

	result := &api.Result{}
	err = stream.RecvMsg(result)
	if err != nil {
		logger.Errorf("Auth Interceptor stream recvMsg failed:%v", err)
		return
	}

	if result.Code != 0 {
		c.JSON(401, "")
		return
	}
	c.Next()
}
