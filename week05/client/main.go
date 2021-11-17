package main

import (
	c "golangwork/week05"
	mc "golangwork/week05/errorCounter"
	"math/rand"
	"sync"
	"time"
)

func main() {
	locker := sync.RWMutex{}
	eCounter := mc.MyCounter{
		Errors: &c.Counter{
			Locker:  &locker,
			Buckets: make(map[int64]*c.BucketValue),
		},
	}

	for i := 0; i < 4; i++ {
		go func() {
			for {
				eCounter.Update(rand.Float64() * 1)
				time.Sleep(2000)
			}
		}()
	}

	select {}
}
