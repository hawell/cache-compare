package cacher

import (
	"github.com/allegro/bigcache"
	"strconv"
	"time"
)

type BigCache struct {
	cache *bigcache.BigCache
	miss int
	hit int
}

func (c *BigCache) Init(cacheSize int, itemSize int) {
	config := bigcache.Config{
		Shards:               1024,
		LifeWindow:           time.Minute,
		CleanWindow:          0,
		MaxEntriesInWindow:   cacheSize,
		MaxEntrySize:         500,
		StatsEnabled:         false,
		Verbose:              false,
		Hasher:               nil,
		HardMaxCacheSize:     (cacheSize * itemSize) / (1024*1024),
	}
	c.cache, _ = bigcache.NewBigCache(config)
}

func (c *BigCache) Get(key interface{}) (interface{}, bool) {
	keyStr := strconv.FormatUint(key.(uint64), 10)
	res, err := c.cache.Get(keyStr)
	if err == bigcache.ErrEntryNotFound {
		c.miss++
		return nil, false
	}
	c.hit++
	return res, true
}

func (c *BigCache) Set(key interface{}, value interface{}) {
	keyStr := strconv.FormatUint(key.(uint64), 10)
	_ = c.cache.Set(keyStr, value.([]byte))
}

func (c *BigCache) HitRatio() float64 {
	return float64(c.hit) / (float64(c.hit) + float64(c.miss))
}
