package handler

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/im/common/api"
	"github.com/im/http/rpc"
	"net/http"
)
import "github.com/im/http/router"

type userHandler struct{}

func (p *userHandler) login(c *gin.Context) {
	req := &api.UserLoginReq{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logger.Errorf("UserHandler login error:%v", err)
		return
	}
	result, err := rpc.UserServiceClientImpl.Login(c, req)
	if err != nil {
		logger.Errorf("UserHandler login error:%v", err)
		return
	}

	res, err := api.ConvertPBResultToJsonResult(result)
	if err != nil {
		logger.Errorf("UserHandler login error:%v", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (p *userHandler) register(c *gin.Context) {
	req := &api.UserRegisterReq{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logger.Errorf("UserHandler register error:%v", err)
		return
	}

	result, err := rpc.UserServiceClientImpl.Register(c, req)
	if err != nil {
		logger.Errorf("UserHandler login error:%v", err)
		return
	}

	res, err := api.ConvertPBResultToJsonResult(result)
	if err != nil {
		logger.Errorf("UserHandler login error:%v", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func init() {
	handler := &userHandler{}
	router.Handlers["/im/user/login"] = handler.login
	router.Handlers["/im/user/register"] = handler.register
}
