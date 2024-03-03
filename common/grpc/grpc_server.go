package grpc

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"fmt"
	"github.com/go-zookeeper/zk"
	"github.com/im/common/constants"
	"github.com/im/common/util"
	pb "google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type GrpcServer struct {
	port             int
	server           *pb.Server
	zkConn           *zk.Conn
	dieChan          chan bool
	heartBeatTimeout time.Duration
	lastAt           int64
}

func NewGrpcServer(port int, srv *pb.Server) *GrpcServer {
	return &GrpcServer{port: port, server: srv, heartBeatTimeout: time.Duration(3) * time.Second}
}

func (p *GrpcServer) Init() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", p.port))
	if err != nil {
		logger.Errorf("GrpcServer Init net listen port:%d failed:%v", p.port, err)
		return err
	}

	addr := config.GetRootConfig().Registries["demoZK"].Address
	conn, _, err := zk.Connect([]string{addr}, time.Duration(3)*time.Second)
	if err != nil {
		logger.Errorf("GrpcServer Init zk connect to %s failed:%v", addr, err)
		return err
	}
	p.zkConn = conn
	b, _, err := p.zkConn.Exists(constants.PushServerPath)
	if err != nil {
		logger.Errorf("GrpcServer Init zk exists %s failed:%v", constants.PushServerPath, err)
		return err
	}

	if !b {
		_, err = p.zkConn.Create(constants.PushServerPath, []byte{}, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			logger.Errorf("GrpcServer Init create node %s failed:%v", constants.PushServerPath, err)
			return err
		}
	}

	nodeName := fmt.Sprintf("%s/%s:%d", constants.PushServerPath, util.NetUtil.Ipv4(), p.port)
	b, stat, err := p.zkConn.Exists(nodeName)
	if err != nil {
		logger.Errorf("GrpcServer Init zk get %s failed:%v", constants.PushServerPath, err)
		return err
	}

	if b {
		err = p.zkConn.Delete(nodeName, stat.Version)
		if err != nil {
			logger.Errorf("GrpcServer Init delete node %s failed:%v", nodeName, err)
			return err
		}

	}
	_, err = p.zkConn.Create(nodeName, p.genData(), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		logger.Errorf("GrpcServer Init Set node %s failed:%v", nodeName, err)
		return err
	}

	go p.heartBeat()

	go p.deleteExpireNode()

	go func() {
		err = p.server.Serve(lis)
		if err != nil {
			logger.Errorf("GrpcServer Init net listen port:%d serve failed:%v", p.port, err)
		}
	}()
	return err
}

func (p *GrpcServer) genData() []byte {
	return []byte(fmt.Sprintf("%d", time.Now().UnixMilli()))
}

func (p *GrpcServer) heartBeat() {
	ticker := time.NewTicker(p.heartBeatTimeout)
	defer func() {
		ticker.Stop()
	}()

	path1 := fmt.Sprintf("%s/%s:%d", constants.PushServerPath, util.NetUtil.Ipv4(), p.port)
	atomic.StoreInt64(&p.lastAt, time.Now().UnixMilli())
	for {
		select {
		case <-ticker.C:
			deadline := time.Now().Add(-2 * p.heartBeatTimeout).UnixMilli()
			if atomic.LoadInt64(&p.lastAt) < deadline {
				logger.Warnf("GrpcServer heartBeat timeout")
				return
			}

			_, state, err := p.zkConn.Get(path1)
			if err != nil {
				logger.Errorf("GrpcServer heartBeat failed:%v", err)
				return
			}
			_, err = p.zkConn.Set(path1, p.genData(), state.Version)
			if err != nil {
				logger.Errorf("GrpcServer heartBeat failed:%v", err)
				return
			}
			//logger.Infof("grpc server heartbeat success...")
			atomic.StoreInt64(&p.lastAt, time.Now().UnixMilli())
		case <-p.dieChan:
			return
		}
	}
}

func (p *GrpcServer) deleteExpireNode() {
	ticker := time.NewTicker(p.heartBeatTimeout)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <-ticker.C:

			children, _, err := p.zkConn.Children(constants.PushServerPath)
			if err != nil {
				logger.Errorf("GrpcServer deleteExpireNode zk get children failed:%v", err)
				return
			}

			for _, child := range children {
				if len(strings.Split(child, ":")) != 2 {
					continue
				}
				path1 := fmt.Sprintf("%s/%s", constants.PushServerPath, child)

				//logger.Warnf("length:%d,%s", len(children), path1)
				bytes, s, err := p.zkConn.Get(path1)
				if err != nil {
					logger.Errorf("GrpcServer deleteExpiredNode zk get:%s failed:%v", path1, err)
					return
				}

				lastAt, err := p.fromData(bytes)
				if err != nil {
					logger.Errorf("GrpcServer deleteExpiredNode fromData:%s failed:%v", string(bytes), err)
					return
				}

				deadline := time.Now().Add(-3 * p.heartBeatTimeout).UnixMilli()
				if lastAt < deadline {
					logger.Warnf("GrpcServer path:%s has expired will be delete from zk", path1)
					err := p.zkConn.Delete(path1, s.Version)
					if err != nil {
						logger.Errorf("GrpcServer delete path %s failed:%v", path1, err)
						return
					}
				}
			}
		case <-p.dieChan:
			return
		}
	}
}

func (p *GrpcServer) fromData(bytes []byte) (int64, error) {
	v, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		logger.Errorf("GrpcServer fromData %s failed:%v", string(bytes), err)
		return 0, err
	}

	return v, nil
}

func (p *GrpcServer) Close() {
	p.zkConn.Close()
	logger.Warnf("GrpcServer closed success...")
}
