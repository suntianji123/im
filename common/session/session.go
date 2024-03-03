package session

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/im/common/network_entity"
	"net"
	"sync"
	"sync/atomic"
)

type Session struct {
	id         int64
	entity     network_entity.NetworkEntity
	onlineInfo *OnlineInfo
	pool       *SessionPool
}

type SessionPool struct {
	sessionByUIDAndAppID  sync.Map
	sessionByID           sync.Map
	sessionIDSrv          *sessionIDService
	sessionCount          int64
	sessionCloseCallbacks []func()
}

type sessionIDService struct {
	sid int64
}

type OnlineInfo struct {
	Uid       int64
	AppId     int
	DeviceId  string
	LoginIp   string
	LoginPort int
	LoginTs   int64
	LoginUuid string
	SocketId  string
}

func NewSessionPool() *SessionPool {
	return &SessionPool{
		sessionIDSrv:         newSessionIDService(),
		sessionByUIDAndAppID: sync.Map{},
		sessionByID:          sync.Map{},
	}
}

func newSessionIDService() *sessionIDService {
	return &sessionIDService{
		sid: 0,
	}
}

func (p *sessionIDService) sessionID() int64 {
	return atomic.AddInt64(&p.sid, 1)
}

func (p *SessionPool) NewSession(entity network_entity.NetworkEntity) *Session {
	s := &Session{
		id:     p.sessionIDSrv.sessionID(),
		entity: entity,
		pool:   p,
	}
	p.sessionByID.Store(s.id, s)
	atomic.AddInt64(&p.sessionCount, 1)
	return s
}

func (p *Session) GetUidAndAppIdKey(info *OnlineInfo) string {
	return fmt.Sprintf("%d:%d", info.Uid, info.AppId)
}

func (p *Session) BindSession(info *OnlineInfo) *Session {
	key := p.GetUidAndAppIdKey(info)
	m, ok := p.pool.sessionByUIDAndAppID.Load(key)
	if !ok {
		m = sync.Map{}
	}
	a := m.(sync.Map)

	pre, _ := a.Load(info.DeviceId)

	a.Store(info.DeviceId, p)
	p.pool.sessionByUIDAndAppID.Store(key, a)
	if pre != nil {
		return pre.(*Session)
	} else {
		return nil
	}
}

func (p *SessionPool) GetSession(uid int64, appId int, deviceId string) *Session {
	key := fmt.Sprintf("%d:%d", uid, appId)
	v, ok := p.sessionByUIDAndAppID.Load(key)
	if ok {
		a := v.(sync.Map)
		if v1, ok1 := a.Load(deviceId); ok1 {
			return v1.(*Session)
		}
	}
	return nil
}

func (p *Session) SetOnlineInfo(info *OnlineInfo) {
	p.onlineInfo = info
}

func (p *Session) GetOnlineInfo() *OnlineInfo {
	return p.onlineInfo
}

func (p *Session) Push(ctx context.Context, message proto.Message) error {
	err := p.entity.Send(ctx, message)
	if err != nil {
		logger.Errorf("Session push failed message:%v,error:%v", message, err)
		return err
	}
	logger.Infof("send msg to session success...")
	return nil
}

func (p *Session) RemoteAddr() net.Addr {
	return p.entity.RemoteAddr()
}

func (p *Session) ID() int64 {
	return p.id
}

func (p *Session) GetAndSetOnlineInfo(info *OnlineInfo) *OnlineInfo {
	old := p.onlineInfo
	p.onlineInfo = info
	return old
}

func (p *Session) IsActive() bool {
	return p.entity.IsActive()
}

func (p *Session) Close() error {
	atomic.AddInt64(&p.pool.sessionCount, -1)
	p.pool.sessionByID.Delete(p.id)

	if p.onlineInfo != nil {
		key := p.GetUidAndAppIdKey(p.onlineInfo)
		v, ok := p.pool.sessionByUIDAndAppID.Load(key)
		if ok {
			m := v.(sync.Map)
			if _, ok = m.Load(p.onlineInfo.DeviceId); ok {
				m.Delete(p.onlineInfo.DeviceId)
			}
		}
	}
	err := p.entity.Close()
	if err != nil {
		logger.Errorf("Session entity failed:%v", err)
		return err
	}
	return nil
}
