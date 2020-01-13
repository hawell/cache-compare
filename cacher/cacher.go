package cacher

const (
	numElements = 2000
	elementSize = 5000
	cacheSize = 500
)

type Cacher interface {
	Init(cacheSize int, itemSize int)
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{})
	HitRatio() float64
}