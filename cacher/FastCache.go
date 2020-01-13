package cacher

import (
	"encoding/binary"
	"github.com/fastcache"
)

type FastCache struct {
	cache *fastcache.Cache
	miss int
	hit int
}

func (c *FastCache) Init(cacheSize int, itemSize int) {
	c.cache = fastcache.New(cacheSize*itemSize)
}

func (c *FastCache) Get(key interface{}) (interface{}, bool) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, key.(uint64))
	res, found := c.cache.HasGet(nil, b)
	if found {
		c.hit++
	} else {
		c.miss++
	}
	return res, found
}
func (c *FastCache) Set(key interface{}, value interface{}) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, key.(uint64))
	c.cache.Set(b, value.([]byte))
}

func (c *FastCache) HitRatio() float64 {
	return float64(c.hit) / (float64(c.hit) + float64(c.miss))
}
