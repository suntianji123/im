package uuid

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"errors"
	"fmt"
	"github.com/go-zookeeper/zk"
	"github.com/im/common/util"
	"math/rand"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	path = "/im/snowflake/nodes"
)

type IdGenerator struct {
	addr             string
	conn             *zk.Conn
	heartBeatTimeout time.Duration
	dieChan          chan bool
	lastAt           int64
	workId           int
	nodeName         string
	generator        *generator
}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{
		heartBeatTimeout: time.Duration(3) * time.Second,
	}
}

func (p *IdGenerator) Init() error {
	addr := config.GetRootConfig().Registries["demoZK"].Address
	conn, _, err := zk.Connect([]string{addr}, time.Duration(3)*time.Second)
	if err != nil {
		logger.Errorf("IdGenerator Init zk connect to %s failed:%v", addr, err)
		return err
	}

	p.conn = conn
	b, _, err := p.conn.Exists(path)
	if err != nil {
		logger.Errorf("IdGenerator Init zk exists path:%s failed:%v", path, err)
		return err
	}

	nodeName := fmt.Sprintf("%s:%d", util.NetUtil.Ipv4(), 1)
	workId := -1
	if b {
		children, _, err := p.conn.Children(path)
		if err != nil {
			logger.Errorf("IdGenerator Init zk Children failed:%v", err)
			return err
		}
		existsMap, err := p.existsMap(children)
		if err != nil {
			logger.Errorf("IdGenerator Init zk existsMap failed:%v", err)
			return err
		}

		workId, err = p.getWorkId(nodeName, existsMap)
		if err != nil {
			logger.Errorf("IdGenerator Init getWorkId failed:%v", err)
			return err
		}
	} else {
		workId = p.findAvailableWorkerId(map[string]int{})
	}

	if workId < 0 || int64(workId) > maxWorkerId {
		return errors.New("IdGenerator generator workId failed with unknow reson")
	}

	p.generator = newGenerator(workId)
	p.workId = workId
	p.nodeName = nodeName
	go p.heartBeat()

	logger.Infof("IdGenerator Init success nodeName:%s,workId:%d", p.nodeName, p.workId)
	return nil
}

func (p *IdGenerator) NextId() (int64, error) {
	v, err := p.generator.nextId()
	if err != nil {
		logger.Errorf("IdGenerator NextId failed:%v", err)
		return 0, err
	}
	return v, nil
}

func (p *IdGenerator) heartBeat() {
	ticker := time.NewTicker(p.heartBeatTimeout)
	defer func() {
		ticker.Stop()
	}()

	atomic.StoreInt64(&p.lastAt, time.Now().UnixMilli())
	path1 := fmt.Sprintf("%s/%s#%d", path, p.nodeName, p.workId)
	for {
		select {
		case <-ticker.C:
			deadline := time.Now().Add(-2 * p.heartBeatTimeout).Unix()
			if atomic.LoadInt64(&p.lastAt) < deadline {
				logger.Warnf("IdGenerator heartBeat timeout")
				return
			}

			_, state, err := p.conn.Get(path1)
			if err != nil {
				logger.Errorf("IdGenerator heartBeat failed:%v", err)
				return
			}
			_, err = p.conn.Set(path1, []byte(p.genData()), state.Version)
			if err != nil {
				logger.Errorf("IdGenerator heartBeat failed:%v", err)
				return
			}
			atomic.StoreInt64(&p.lastAt, time.Now().UnixMilli())
		case <-p.dieChan:
			return
		}
	}
}

func (p *IdGenerator) deleteExpireNode() {
	ticker := time.NewTicker(p.heartBeatTimeout)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <-ticker.C:
			children, _, err := p.conn.Children(path)
			if err != nil {
				logger.Errorf("IdGenerator deleteExpireNode zk get children failed:%v", err)
				return
			}

			for _, child := range children {
				path1 := fmt.Sprintf("%s/%s", path, child)
				bytes, s, err := p.conn.Get(path1)
				if err != nil {
					logger.Errorf("IdGenerator deleteExpiredNode zk get:%s failed:%v", path1, err)
					return
				}

				lastAt, err := p.fromData(bytes)
				if err != nil {
					logger.Errorf("IdGenerator deleteExpiredNode fromData:%s failed:%v", string(bytes), err)
					return
				}

				deadline := time.Now().Add(-24 * time.Hour).Unix()
				if lastAt < deadline {
					logger.Warnf("IdGeneartor path:%s has expired will be delete from zk", path1)
					err := p.conn.Delete(path1, s.Version)
					if err != nil {
						logger.Errorf("IdGenerator delete path %s failed:%v", path1, err)
						return
					}
				}
			}
		case <-p.dieChan:
			return
		}
	}
}

func (p *IdGenerator) fromData(bytes []byte) (int64, error) {
	v, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		logger.Errorf("IdGenerator fromData %s failed:%v", string(bytes), err)
		return 0, err
	}

	return v, nil
}

func (p *IdGenerator) getWorkId(nodeName string, existsMap map[string]int) (int, error) {
	if workId, ok := existsMap[nodeName]; ok {
		path1 := fmt.Sprintf("%s/%s#%d", path, nodeName, workId)
		_, state, err := p.conn.Get(path1)
		if err != nil {
			logger.Errorf("IdGenerator getWorkId get %s#%d failed:%v", nodeName, workId, err)
			return 0, err
		}

		_, err = p.conn.Set(path1, []byte(p.genData()), state.Version)
		if err != nil {
			logger.Errorf("IdGenerator getWorkId set %s#%d failed:%v", nodeName, workId, err)
			return 0, err
		}
		return workId, nil
	}

	workId := p.findAvailableWorkerId(existsMap)
	_, err := p.conn.Create(fmt.Sprintf("%s/%s#%d", path, nodeName, workId), []byte(p.genData()), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		logger.Errorf("IdGenerator getWorkId failed:%v", err)
		return 0, err
	}
	return workId, nil
}

func (p *IdGenerator) findAvailableWorkerId(existsMap map[string]int) int {
	i := rand.Intn(1024)
	if !p.mapExistsValue(existsMap, i) {
		return i
	}

	j := (i + 1) % 1024
	for j != i {
		if !p.mapExistsValue(existsMap, j) {
			break
		}

		j = (j + 1) % 1024
	}
	return j
}

func (p *IdGenerator) mapExistsValue(m map[string]int, val int) bool {
	for _, v := range m {
		if v == val {
			return true
		}
	}

	return false
}

func (p *IdGenerator) genData() string {
	return fmt.Sprintf("%d", time.Now().UnixMilli())
}

func (p *IdGenerator) existsMap(children []string) (map[string]int, error) {
	m := make(map[string]int, 0)
	if children == nil {
		return m, nil
	}

	for _, child := range children {
		name, workId, err := p.parseNode(child)
		if err != nil {
			logger.Errorf("IdGenerator existsMap parseNode %s failed:%v", err)
			return nil, err
		}

		m[name] = workId
	}
	return m, nil
}

func (p *IdGenerator) parseNode(child string) (string, int, error) {
	strs := strings.Split(child, "#")
	workId := -1
	var err error
	if len(strs) > 1 {
		workId, err = strconv.Atoi(strs[1])
		if err != nil {
			logger.Errorf("IdGenerator parseNode strconv Atoi:%s failed:%v", child, err)
			return "", 0, err
		}

		return strs[0], workId, nil
	}
	return "", workId, nil
}

func (p *IdGenerator) Stop() {
	p.conn.Close()
	//	close(p.dieChan)
}
