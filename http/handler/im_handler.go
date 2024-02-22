package handler

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/im/common/api"
	"github.com/im/http/rpc"
	"net/http"
)

type imHandler struct {
}

func (*imHandler) sync(c *gin.Context) {
	req := &api.SyncReq{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logger.Errorf("imHandler list error:%v", err)
		return
	}
	result, err := rpc.SyncServiceClientImpl.Sync(c, req)
	if err != nil {
		logger.Errorf("imHandler list error:%v", err)
		return
	}

	res, err := api.ConvertPBResultToJsonResult(result)
	if err != nil {
		logger.Errorf("imHandler list error:%v", err)
		return
	}

	c.JSON(http.StatusOK, res)
}
