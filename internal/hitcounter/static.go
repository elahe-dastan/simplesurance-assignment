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
	hits          [window]int64
}

func NewStatic() *StaticHitCounter {
	return &StaticHitCounter{
		lock:          sync.Mutex{},
		lastInsertion: 0,
		startupTime:   time.Now().Unix(),
		hits:          [window]int64{},
	}
}

func (hc *StaticHitCounter) Hit(seconds int64) {
	hc.lock.Lock()
	defer hc.lock.Unlock()

	timeFromStart := seconds - hc.startupTime

	if timeFromStart > window {
		// This shows the last cell of previous window
		preWindowEnd := seconds % window
		cellsToRemove := seconds - hc.lastInsertion

		if cellsToRemove > window {
			hc.hits = [window]int64{}
		} else {
			for i := preWindowEnd; i > preWindowEnd-cellsToRemove; i-- {
				hc.hits[i] = 0
			}
		}
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
