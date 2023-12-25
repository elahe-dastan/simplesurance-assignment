package internal

import (
	"sync"
	"time"
)

var upTimestamp = time.Now().Unix()
var window = int64(60)

type HitCounter struct {
	lock          sync.Mutex
	lastInsertion int64
	hits          [60]int64
}

func (hc *HitCounter) Hit(seconds int64) {
	hc.lock.Lock()
	defer hc.lock.Unlock()

	//timestamp := time.Now().Unix()
	diff := seconds - upTimestamp

	if diff > window {
		for i := seconds % window; i > (seconds%window)-(seconds-hc.lastInsertion); i-- {
			hc.hits[i] = 0
		}
	} else {
		hc.hits[seconds] = 0
	}

	hc.lastInsertion = seconds
	hc.hits[seconds%window] = 1
}

func (hc *HitCounter) GetHitCount() int64 {
	res := int64(0)
	for i := int64(0); i < window; i++ {
		res += hc.hits[i]
	}

	return res
}
