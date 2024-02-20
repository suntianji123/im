package router

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/http/rpc"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

var Handlers = map[string]gin.HandlerFunc{}

type Router struct {
	router *gin.Engine
}

type serverConfig struct {
	Port int
}

func NewRouter() *Router {
	return &Router{
		router: gin.Default(),
	}
}

func (r *Router) globalInterceptor(c *gin.Context) {
	path := c.Request.RequestURI
	if path == "/im/user/login" || path == "/im/user/register" {
		c.Next()
		return
	}

	remote, err := rpc.AuthServiceClientImpl.CheckAuth(c)
	if err != nil {
		logger.Errorf("Router globalInterceptor failed:%v", err)
		return
	}

	uid, err := strconv.ParseInt(c.GetHeader("uid"), 10, 64)
	if err != nil {
		logger.Errorf("Router globalInterceptor failed:%v", err)
		return
	}

	appId, err := strconv.ParseInt(c.GetHeader("appId"), 10, 32)
	if err != nil {
		logger.Errorf("Router globalInterceptor failed:%v", err)
		return
	}

	deviceType, err := strconv.ParseInt(c.GetHeader("deviceType"), 10, 32)
	if err != nil {
		logger.Errorf("Router globalInterceptor failed:%v", err)
		return
	}

	err = remote.Send(&api.AuthCheckReq{
		Uid:        uid,
		AppId:      int32(appId),
		DeviceType: int32(deviceType),
		Token:      c.GetHeader("token"),
	})
	if err != nil {
		logger.Errorf("Router globalInterceptor failed:%v", err)
		return
	}

	result := &api.Result{}
	err = remote.RecvMsg(result)
	if err != nil {
		logger.Errorf("Router globalInterceptor failed:%v", err)
		return
	}
	if result.Code != constants.SuccessCode {
		c.JSON(401, "")
		return
	}
	c.Next()
}

func (r *Router) Init() error {

	r.router.Use(r.globalInterceptor)

	for path, fun := range Handlers {
		r.router.POST(path, fun)
	}

	srvConfig := &serverConfig{}
	err := mapstructure.Decode(config.GetRootConfig().Custom.GetDefineValue("server", map[string]interface{}{}), srvConfig)
	if err != nil {
		logger.Errorf("启动http服务器失败:%v", err)
		return err
	}

	err = r.router.Run(fmt.Sprintf(":%d", srvConfig.Port))
	if err != nil {
		logger.Errorf("启动http服务器失败:%v", err)
		return err
	}
	return nil
}
