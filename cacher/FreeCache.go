package cacher

import (
	"encoding/binary"
	"fmt"
	"github.com/coocood/freecache"
)

type FreeCache struct {
	cache *freecache.Cache
	miss int
	hit int
}

func (c *FreeCache) Init(cacheSize int, itemSize int) {
	c.cache = freecache.NewCache(cacheSize*itemSize)
}

func (c *FreeCache) Get(key interface{}) (interface{}, bool) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, key.(uint64))
	res, err := c.cache.Get(b)
	if err == freecache.ErrNotFound {
		c.miss++
		return nil, false
	} else {
		c.hit++
		return res, true
	}
}

func (c *FreeCache) Set(key interface{}, value interface{}) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, key.(uint64))
	if err := c.cache.Set(b, value.([]byte), 60); err != nil {
		fmt.Println(err)
	}
}

func (c *FreeCache) HitRatio() float64 {
	return float64(c.hit) / (float64(c.hit) + float64(c.miss))
}
