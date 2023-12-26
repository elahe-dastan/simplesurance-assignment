package hitcounter

import (
	"encoding/json"
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

// FromFileStatic reads saved HitCounter data from disk and returns a *StaticHitCounter
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

// ToFileStatic locks the HitCounter data structure and saves it to a file
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

// WriteNewRequest updated last insertion time and increase the corresponding cell in Hit array
func (hc *StaticHitCounter) WriteNewRequest(seconds int64) {
	hc.LastInsertion = seconds
	hc.Hits[seconds%window] += 1
}

// Hit locks HitCounter data structure to ensure thread safety. The method processes incoming requests by updating
// the request count in a time-based sliding window. It operates on a fixed-size array 'Hits' of length 'window', where
// each cell represents one second within the window. The method first calculates the time elapsed since the hit
// counter's startup and adjusts for the sliding window:
// 1. If the current request falls into a new window, it resets the expired cells. The reset mechanism depends on the
// time elapsed since the last request:
//
//		 a. If the time difference exceeds the window size, it indicates that all cells are expired. Hence, the entire
//			 'Hits' array is reset to zero.
//		 b. Otherwise, it calculates the number of expired cells based on the time difference and resets them starting
//			 from the end of the previous window.
//	 2. The method then increments the count in the cell corresponding to the current time slot, determined by
//	    'seconds % window'.
//
// If this is the initial window since the counter's startup, no cells are reset.
func (hc *StaticHitCounter) Hit(seconds int64) {
	hc.lock.Lock()
	defer hc.lock.Unlock()

	timeFromStart := seconds - hc.StartupTime
	if timeFromStart < window {
		hc.WriteNewRequest(seconds)
		return
	}

	// This shows the last cell of previous window
	preWindowEnd := seconds % window
	cellsToRemove := seconds - hc.LastInsertion

	if cellsToRemove >= window {
		hc.Hits = [window]int64{}
	} else {
		for i := preWindowEnd; i > preWindowEnd-cellsToRemove; i-- {
			hc.Hits[i] = 0
		}
	}

	hc.WriteNewRequest(seconds)
}

// Count sums up the Hits array and returns the total number of requests in the last window
func (hc *StaticHitCounter) Count() int64 {
	res := int64(0)
	for i := int64(0); i < window; i++ {
		res += hc.Hits[i]
	}

	return res
}
