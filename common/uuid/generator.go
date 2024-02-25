package uuid

import (
	"github.com/im/common/constants"
	"sync"
	"sync/atomic"
	"time"
)

const (
	epoch int64 = 1619619923000

	workerIdBits          int64 = 10
	maxWorkerId           int64 = int64(-1) ^ (int64(-1) << workerIdBits)
	sequenceBits          int64 = 12
	workerIdShift               = sequenceBits
	timestampLeftShiftint       = sequenceBits + workerIdBits
	sequenceMask          int64 = int64(-1) ^ (int64(-1) << sequenceBits)
)

type generator struct {
	workId   int
	lastAt   int64
	sequence int64
	lock     sync.Mutex
}

func newGenerator(workId int) *generator {
	return &generator{workId: workId}
}

func (p *generator) nextId() (int64, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	now := p.timeGen()
	if now < p.lastAt {
		return 0, constants.ErrIdGeneartorTimeCallback
	}

	if now == p.lastAt {
		p.sequence = (p.sequence + 1) & sequenceMask
		if p.sequence == 0 {
			now = p.timeNexMilli(atomic.LoadInt64(&p.lastAt))
		}
	} else {
		p.sequence = 0
	}
	p.lastAt = now
	return ((now - epoch) << timestampLeftShiftint) | (int64(p.workId) << workerIdShift) | p.sequence, nil
}

func (p *generator) timeGen() int64 {
	return time.Now().UnixMilli()
}

func (p *generator) timeNexMilli(lastAt int64) int64 {
	now := p.timeGen()
	for now <= lastAt {
		now = p.timeGen()
	}
	return now
}

func GetTs(msgId int64) int64 {
	return (msgId >> timestampLeftShiftint) + epoch
}

func GenMsgId(ts int64) int64 {
	return (ts - epoch) << timestampLeftShiftint
}
