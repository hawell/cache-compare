package cacher

import "github.com/dgraph-io/ristretto"

type Ristretto struct {
	ristrettoCache *ristretto.Cache
	hit int
	miss int
}

func (c *Ristretto) Init(cacheSize int, itemSize int) {
	c.ristrettoCache, _ = ristretto.NewCache(&ristretto.Config{
		NumCounters: 10 * int64(cacheSize),
		MaxCost:     int64(cacheSize),
		BufferItems: 64,
	})
}

func (c *Ristretto) Get(key interface{}) (interface{}, bool) {
	res, found := c.ristrettoCache.Get(key)
	if found {
		c.hit++
	} else {
		c.miss++
	}
	return res, found
}

func (c *Ristretto) Set(key interface{}, value interface{}) {
	c.ristrettoCache.Set(key, value, 1)
}

func (c *Ristretto) HitRatio() float64 {
	return float64(c.hit) / (float64(c.hit) + float64(c.miss))
}
