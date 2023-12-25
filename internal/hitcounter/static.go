package hitcounter

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

const window = int64(60)

type StaticHitCounter struct {
	lock          sync.Mutex
	LastInsertion int64         `json:"last_insertion,omitempty"`
	StartupTime   int64         `json:"startup_time,omitempty"`
	Hits          [window]int64 `json:"hits,omitempty"`
}

func NewStatic() *StaticHitCounter {
	return &StaticHitCounter{
		lock:          sync.Mutex{},
		LastInsertion: 0,
		StartupTime:   time.Now().Unix(),
		Hits:          [window]int64{},
	}
}

func FromFileStatic(fn string) (*StaticHitCounter, error) {
	bytes, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	hc := new(StaticHitCounter)

	if err := json.Unmarshal(bytes, hc); err != nil {
		return nil, err
	}

	return hc, nil
}

func ToFileStatic(fn string, hc *StaticHitCounter) error {
	hc.lock.Lock()
	defer hc.lock.Unlock()

	bytes, err := json.Marshal(hc)
	if err != nil {
		return err
	}

	if err := os.WriteFile(fn, bytes, 0644); err != nil {
		return err
	}

	return nil
}

func (hc *StaticHitCounter) Hit(seconds int64) {
	hc.lock.Lock()
	defer hc.lock.Unlock()

	timeFromStart := seconds - hc.StartupTime

	if timeFromStart >= window {
		// This shows the last cell of previous window
		preWindowEnd := seconds % window
		cellsToRemove := seconds - hc.LastInsertion

		for i := preWindowEnd; i > preWindowEnd-cellsToRemove; i-- {
			if i < 0 {
				fmt.Println(i)
				fmt.Println(((i % window) + window) % window)
				hc.Hits[((i%window)+window)%window] = 0
			} else {
				hc.Hits[i] = 0
			}
		}
	}

	hc.LastInsertion = seconds
	hc.Hits[seconds%window] += 1
}

func (hc *StaticHitCounter) Count() int64 {
	res := int64(0)
	for i := int64(0); i < window; i++ {
		res += hc.Hits[i]
	}

	return res
}
