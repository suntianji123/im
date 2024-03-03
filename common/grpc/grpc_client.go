package grpc

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/common/api"
	"github.com/shimingyah/pool"
	pb "google.golang.org/grpc"
	"sync/atomic"
	"time"
)

type GrpcClient struct {
	addr              string
	pool              pool.Pool
	heartBeatTimeout  time.Duration
	dieChan           chan bool
	lastAt            int64
	running           bool
	heartBeartService api.HeartBeartServiceClient
}

func NewGrpcClient(addr string) *GrpcClient {
	return &GrpcClient{
		addr:             addr,
		heartBeatTimeout: time.Duration(2) * time.Second,
		dieChan:          make(chan bool),
	}
}

func (g *GrpcClient) Init() error {
	p, err := pool.New(g.addr, pool.DefaultOptions)
	if err != nil {
		logger.Errorf("GppcClient addr:%s Init failed:%v", g.addr, err)
		return err
	}

	conn, err := p.Get()
	if err != nil {
		logger.Errorf("GrpcClient addr:%s get conn failed:%v", g.addr, err)
		return err
	}

	g.running = true
	g.pool = p
	g.heartBeartService = api.NewHeartBeartServiceClient(conn.Value())
	go g.sendHeartBeart()
	go g.heartBeat()
	logger.Infof("grpcClient:%s Init success...", g.addr)
	return nil
}

func (g *GrpcClient) sendHeartBeart() {
	ticker := time.NewTicker(g.heartBeatTimeout)
	defer func() {
		ticker.Stop()
		logger.Warnf("grpcClient addr:%s sendHeart close...", g.addr)
		g.Close()
	}()
	for {
		select {
		case <-ticker.C:
			_, err := g.heartBeartService.HeartBeart(context.Background(), &api.PPing{})
			if err != nil {
				logger.Errorf("grpcClient sendHeartBeat failed:%v", err)
				return
			}
			atomic.StoreInt64(&g.lastAt, time.Now().UnixMilli())
		case <-g.dieChan:
			logger.Warnf("grpcClient:%s closed", g.addr)
			return
		}
	}
}

func (g *GrpcClient) heartBeat() {
	ticker := time.NewTicker(g.heartBeatTimeout)
	defer func() {
		ticker.Stop()
		g.Close()
		logger.Warnf("grpcClient addr:%s heartbeart close...", g.addr)
	}()

	atomic.StoreInt64(&g.lastAt, time.Now().UnixMilli())
	for g.running {
		select {
		case <-ticker.C:
			deadline := time.Now().Add(-2 * g.heartBeatTimeout).UnixMilli()
			if atomic.LoadInt64(&g.lastAt) < deadline {
				logger.Debugf("grpcClient:%s heartBeat timeout", g.addr)
				return
			}
		case <-g.dieChan:
			logger.Warnf("grpcClient:%s closed", g.addr)
			return
		}
	}
}

func (g *GrpcClient) SetLastAl() {
	atomic.StoreInt64(&g.lastAt, time.Now().UnixMilli())
}

func (g *GrpcClient) SetStop() {
	g.dieChan <- true
}

func (g *GrpcClient) Close() {
	g.running = false
	g.pool.Close()
}

func (g *GrpcClient) Running() bool {
	return g.running
}

func (g *GrpcClient) GetConn() (*pb.ClientConn, error) {
	c, err := g.pool.Get()
	if err != nil {
		logger.Errorf("GrpcClient %s GetConn failed:%v", g.addr, err)
		return nil, err
	}
	return c.Value(), nil
}
