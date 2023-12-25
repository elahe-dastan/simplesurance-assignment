package internal

import (
	"encoding/gob"
	"os"
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

func (hc *HitCounter) Hit(seconds int64, filename string) {
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

// Save method to serialize and save the HitCounter to a file
func (hc *HitCounter) Save(filename string) error {
	hc.lock.Lock()
	defer hc.lock.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(hc)
	if err != nil {
		return err
	}

	return nil
}

// Load method to load and deserialize the HitCounter from a file
func Load(filename string) (*HitCounter, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hc := new(HitCounter)
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(hc)
	if err != nil {
		return nil, err
	}

	return hc, nil
}
