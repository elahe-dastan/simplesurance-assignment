package hitcounter

import (
	"sync"
	"time"
)

const window = int64(60)

type StaticHitCounter struct {
	lock          sync.Mutex
	lastInsertion int64
	startupTime   int64
	hits          [60]int64
}

func NewStatic() *StaticHitCounter {
	return &StaticHitCounter{
		lock:          sync.Mutex{},
		lastInsertion: 0,
		startupTime:   time.Now().Unix(),
		hits:          [60]int64{},
	}
}

func (hc *StaticHitCounter) Hit(seconds int64) {
	hc.lock.Lock()
	defer hc.lock.Unlock()

	diff := seconds - hc.startupTime

	if diff > window {
		for i := seconds % window; i > (seconds%window)-(seconds-hc.lastInsertion); i-- {
			hc.hits[i] = 0
		}
	} else {
		hc.hits[seconds] = 0
	}

	hc.lastInsertion = seconds
	hc.hits[seconds%window] += 1
}

func (hc *StaticHitCounter) Count() int64 {
	res := int64(0)
	for i := int64(0); i < window; i++ {
		res += hc.hits[i]
	}

	return res
}
