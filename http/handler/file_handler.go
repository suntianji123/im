package handler

import (
	"bytes"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/im/common/api"
	"github.com/im/common/conf"
	"github.com/im/http/app"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"net/http"
)

type fileHandler struct {
	config *conf.QiNiuYunConfig
}

func (p *fileHandler) Init() {
	p.config = app.App.GetBuilder().GetQiNiuYunConfig()
}

func (p *fileHandler) upload(c *gin.Context) {
	f, err := c.FormFile("f")
	if err != nil {
		logger.Errorf("fileHandler upload failed:%v", err)
		return
	}

	putPolicy := storage.PutPolicy{
		Scope: p.config.Bucket,
	}
	mac := auth.New(p.config.AccessKey, p.config.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := &storage.PutRet{}

	key := f.Filename
	f1, err := f.Open()
	defer f1.Close()
	if err != nil {
		logger.Errorf("fileHandler upload failed:%v", err)
		return
	}

	buff := new(bytes.Buffer)
	dataLen, err := io.Copy(buff, f1)
	if err != nil {
		logger.Errorf("fileHandler upload failed:%v", err)
		return
	}
	err = formUploader.Put(c, ret, upToken, key, buff, dataLen, &storage.PutExtra{})
	if err != nil {
		logger.Errorf("fileHandler upload failed:%v", err)
		return
	}

	res := &api.JsonResult{Data: &api.FileUploadResp{
		//Url: strings.ReplaceAll(fmt.Sprintf("http://%s/%s", p.config.Domain, ret.Key), " ", ""),
		Url: fmt.Sprintf("http://%s/%s", p.config.Domain, ret.Key),
	}}

	c.JSON(http.StatusOK, res)
}
