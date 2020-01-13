package cacher

import (
	"github.com/karlseguin/ccache"
	"strconv"
	"time"
)

type CCache struct {
	cache *ccache.Cache
	miss int
	hit int
}

func (c *CCache) Init(cacheSize int, itemSize int) {
	c.cache = ccache.New(ccache.Configure().MaxSize(int64(cacheSize)))
}

func (c *CCache) Get(key interface{}) (interface{}, bool) {
	keyStr := strconv.FormatUint(key.(uint64), 10)
	res := c.cache.Get(keyStr)
	if res == nil {
		c.miss++
		return nil, false
	} else {
		c.hit++
		return res.Value(), true
	}
}

func (c *CCache) Set(key interface{}, value interface{}) {
	keyStr := strconv.FormatUint(key.(uint64), 10)
	c.cache.Set(keyStr, value, time.Minute*10)
}

func (c *CCache) HitRatio() float64 {
	return float64(c.hit) / (float64(c.hit) + float64(c.miss))
}
