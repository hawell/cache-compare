package main

import (
	"fmt"
	"github.com/hawell/cache-compare/cacher"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// TestHitRatio()
	TestThroughput()
}

func TestHitRatio() {
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
	ratio := 0.01
	cacheSize := 60 * 1024
	itemSize := 1024
	dbSize := int(float64(cacheSize) / ratio)
	queryCount := 10000000

	for i := range caches {
		caches[i].Init(cacheSize, itemSize)
	}

	entry := make([]byte, itemSize)
	var wg sync.WaitGroup

	for _, cache := range caches {
		go func(cache cacher.Cacher) {
			wg.Add(1)
			source := rand.NewSource(rand.Int63n(int64(time.Now().Nanosecond())))
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

func TestThroughput() {
	read := func(cache cacher.Cacher, zipf *rand.Zipf, entry []byte) {
		z := zipf.Uint64()
		// log.Println(z)
		_, found := cache.Get(z)
		if !found {
			cache.Set(z, entry)
		}
	}
	write := func(cache cacher.Cacher, zipf *rand.Zipf, entry []byte) {
		z := zipf.Uint64()
		cache.Set(z, entry)
	}

	caches := []cacher.Cacher{
		&cacher.GoBurrow{},
		&cacher.Ristretto{},
		&cacher.CCache{},
		&cacher.FastCache{},
		&cacher.FreeCache{},
		&cacher.BigCache{},
	}

	// ratio 1 = all read
	// ratio 0 = all write
	ratio := 0.0
	goroutines := []int{4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48, 52, 56, 60, 64}
	cacheSize := 60 * 1024
	itemSize := 1024
	dbSize := cacheSize
	duration := 5

	for _, cache := range caches {
		cache.Init(cacheSize, itemSize)
	}

	entry := make([]byte, itemSize)
	for _, cache := range caches {
		var wg sync.WaitGroup
		for _, count := range goroutines {
			sum := make([]int, count)
			for i := 0; i < count; i++ {
				var fn func(cacher.Cacher, *rand.Zipf, []byte)
				if i < int(float64(count)*ratio) {
					fn = read
				} else {
					fn = write
				}
				go func(i int, fn func(cacher.Cacher, *rand.Zipf, []byte)) {
					wg.Add(1)
					source := rand.NewSource(rand.Int63n(int64(time.Now().Nanosecond())))
					r := rand.New(source)
					zipf := rand.NewZipf(r, 1.00001, 1, uint64(dbSize))
					t := time.NewTicker(time.Second * time.Duration(duration))
					for {
						select {
						case <-t.C:
							wg.Done()
							return
						default:
							fn(cache, zipf, entry)
							sum[i]++
						}
					}
				}(i, fn)
			}
			time.Sleep(time.Second)
			wg.Wait()
			total := 0
			for i := range sum {
				total += sum[i]
			}
			fmt.Print(total/duration, ",")
		}
		fmt.Println()
	}
}