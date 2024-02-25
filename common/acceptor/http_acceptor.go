package acceptor

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

type HttpAcceptor struct {
	port        int
	router      *gin.Engine
	handlers    map[string]gin.HandlerFunc
	interceptor func(c *gin.Context)
}

func NewHttpAcceptor(port int) *HttpAcceptor {
	return &HttpAcceptor{
		router:   gin.Default(),
		handlers: make(map[string]gin.HandlerFunc),
		port:     port,
	}
}

func (p *HttpAcceptor) RegisterHandlers(handlers map[string]gin.HandlerFunc) {
	p.handlers = handlers
}

func (p *HttpAcceptor) RegisterInterceptor(interceptor gin.HandlerFunc) {
	p.interceptor = interceptor
}

func (p *HttpAcceptor) Init() error {
	if p.interceptor != nil {
		p.router.Use(p.interceptor)
	}

	for path, fun := range p.handlers {
		p.router.POST(path, fun)
	}

	go func() {
		err := p.router.Run(fmt.Sprintf(":%d", p.port))
		logger.Errorf("启动http服务器失败:%v", err)
	}()
	return nil
}

func (p *HttpAcceptor) Stop() error {
	logger.Warnf("close http server success")
	return nil
}
