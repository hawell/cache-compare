package main

import (
	"fmt"
	"github.com/hawell/cache-compare/cacher"
	"math/rand"
	"sync"
	"time"
)

func main() {

	caches := []cacher.Cacher {
		&cacher.Optimal{},
		&cacher.Ristretto{},
		&cacher.BigCache{},
		&cacher.FreeCache{},
		&cacher.FastCache{},
		&cacher.GoBurrow{},
		&cacher.CCache{},
	}

	// Minimum cacheSize*itemSize for FastCache is 32MB
	// itemSize should be less than 1/1024 of cacheSize*itemSize for FreeCache
	cacheSize := 60 * 1024
	itemSize := 1024
	dbSize := cacheSize
	queryCount := 10000000

	for i := range caches {
		caches[i].Init(cacheSize, itemSize)
	}

	entry := make([]byte, itemSize)
	var wg sync.WaitGroup

	for _, cache := range caches {
		go func(cache cacher.Cacher) {
			wg.Add(1)
			source := rand.NewSource(rand.Int63n(30))
			r := rand.New(source)
			zipf := rand.NewZipf(r, 1.00001, 1, uint64(dbSize))
			for i:=0 ;i < queryCount; i++ {
				z := zipf.Uint64()
				// log.Println(z)
				_, found := cache.Get(z)
				if !found {
					cache.Set(z, entry)
				}
			}
			wg.Done()
		}(cache)
	}
	time.Sleep(time.Second)

	wg.Wait()

	for i := range caches {
		fmt.Println(caches[i].HitRatio())
	}
}
