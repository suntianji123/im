package rpc

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"errors"
	"fmt"
	"github.com/dubbogo/go-zookeeper/zk"
	"github.com/im/common/api"
	"github.com/im/common/constants"
	"github.com/im/common/grpc"
	"sync"
	"time"
)

var MessageServiceClientImpl = &api.MessageServiceClientImpl{}

var BatchServiceClientImpl = &batchServiceClient{
	childrenChan: make(chan []string),
	dieChan:      make(chan bool),
}

type batchServiceClient struct {
	pathCache    sync.Map
	zkConn       *zk.Conn
	dieChan      chan bool
	running      bool
	childrenChan chan []string
}

type childrenEvent struct {
	children []string
	event    *zk.Watcher
}

func init() {
	config.SetConsumerService(MessageServiceClientImpl)
}

func (p *batchServiceClient) Init() error {
	addr := config.GetRootConfig().Registries["demoZK"].Address
	conn, _, err := zk.Connect([]string{addr}, time.Duration(3)*time.Second)
	if err != nil {
		logger.Errorf("batchServiceClient Init zk connect to %s failed:%v", addr, err)
		return err
	}

	p.running = true
	p.zkConn = conn

	go p.callBack()
	go p.watch(conn)

	return nil
}

func (g *batchServiceClient) Batch(path string, batch *api.PBatch) error {
	if v, ok := g.pathCache.Load(path); ok {
		c := v.(*grpc.GrpcClient)
		if c.Running() {
			conn, err := c.GetConn()
			if err != nil {
				logger.Errorf("batchServiceClient Batch failed:%v", err)
				return err
			}
			_, err = api.NewBatchServiceClient(conn).Batch(context.Background(), batch)
			if err != nil {
				logger.Errorf("batchServiceClient Batch failed:%v", err)
				return err
			}
			return nil
		}
	}

	logger.Errorf("grpcClient path %s is Closed,batch failed", path)
	return errors.New(fmt.Sprintf("grpcClient path %s is Closed,batch failed", path))
}

func (g *batchServiceClient) callBack() {
	defer close(g.childrenChan)

	for {
		select {
		case children := <-g.childrenChan:
			//logger.Warnf("当前连接:%v", children)
			for _, child := range children {
				a, ok := g.pathCache.Load(child)
				//logger.Warnf("当前连接：%v,存在：%v", child, ok)
				if ok {
					c := a.(*grpc.GrpcClient)
					if c.Running() {
						c.SetStop()
					}

					err := c.Init()
					if err != nil {
						logger.Errorf("batchServiceClient callBack path %s failed:%v", child, err)
						return
					}
				} else {
					c := grpc.NewGrpcClient(child)
					if err := c.Init(); err != nil {
						logger.Errorf("batchServicClient path %s grpcClient init failed:%v", child, err)
						return
					}

					g.pathCache.Store(child, c)
				}

			}

		case <-g.dieChan:
			return
		}
	}
}

func (g *batchServiceClient) watch(conn *zk.Conn) error {

	b, _, _, err := conn.ExistsW(constants.PushServerPath)
	if err != nil {
		logger.Errorf("batchServiceClient zk exists watch %s failed:%v", constants.PushServerPath, err)
		return err
	}

	if !b {
		_, err = conn.Create(constants.PushServerPath, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			logger.Errorf("batchServiceClient zk create path %s failed:%v", constants.PushServerPath, err)
			return err
		}
	}

	for g.running {
		children, _, event, err := conn.ChildrenW(constants.PushServerPath)
		logger.Warnf("children changed:%v", children)
		if err != nil {
			logger.Errorf("batchServiceClient zk ChildrenW watch %s failed:%v", constants.PushServerPath, err)
			return err
		}
		g.childrenChan <- children

		select {
		case <-event.EvtCh:
		}

	}

	return nil
}

func (g *batchServiceClient) onNodeCreated(path string) error {
	grpcClient := grpc.NewGrpcClient(path)
	err := grpcClient.Init()
	if err != nil {
		logger.Errorf("batchServiceClient onNodeCreated path %s failed:%v", path, err)
		return err
	}
	g.pathCache.Store(path, grpcClient)
	return nil
}

func (g *batchServiceClient) onNodeDeleted(path string) error {
	v, b := g.pathCache.LoadAndDelete(path)
	if b {
		v.(*grpc.GrpcClient).SetStop()
	}
	return nil
}

func (g *batchServiceClient) Close() {
	g.running = false
	g.zkConn.Close()
	logger.Warnf("batchServiceClient close")
}
