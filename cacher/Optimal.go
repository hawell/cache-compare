package cacher

type Optimal struct {
	size int
	cache []interface{}
	hit int
	miss int
}

func (c *Optimal) Init(cacheSize int, itemSize int) {
	c.size = cacheSize
	c.cache = make([]interface{}, cacheSize)
}

func (c *Optimal) Get(key interface{}) (interface{}, bool) {
	i := key.(uint64)
	if i >= uint64(c.size) {
		c.miss++
		return nil, false
	}
	if c.cache[i] == nil {
		c.miss++
		return c.cache[i], false
	} else {
		c.hit++
		return c.cache[i], true
	}
}

func (c *Optimal) Set(key interface{}, value interface{}) {
	i := key.(uint64)
	if i < uint64(c.size) {
		c.cache[i] = value
	}
}

func (c *Optimal) HitRatio() float64 {
	return float64(c.hit) / (float64(c.hit) + float64(c.miss))
}
