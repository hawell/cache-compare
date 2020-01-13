package cacher

import (
	"github.com/goburrow/cache"
	"time"
)

type GoBurrow struct {
	cache cache.Cache
	miss int
	hit int
}

func (c *GoBurrow) Init(cacheSize int, itemSize int) {
	c.cache = cache.New(cache.WithMaximumSize(cacheSize),
		cache.WithExpireAfterAccess(time.Minute*10),
		cache.WithRefreshAfterWrite(time.Minute*10))
}

func (c *GoBurrow) Get(key interface{}) (interface{}, bool) {
	res, found := c.cache.GetIfPresent(key)
	if found {
		c.hit++
	} else {
		c.miss++
	}
	return res, found
}

func (c *GoBurrow) Set(key interface{}, value interface{}) {
	c.cache.Put(key, value)
}

func (c *GoBurrow) HitRatio() float64 {
	return float64(c.hit) / (float64(c.hit) + float64(c.miss))
}
