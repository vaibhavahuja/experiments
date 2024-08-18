package main

// Cache any struct which implements these methods will be my cache
// in case of chained cache I will have 3 layers of caching
// all implementing the same struct
type Cache[K comparable, V any] interface {
	AddKey(K, V) error
	GetVal(K) (V, error)
}

// Enum-like constants for different cache types
const (
	InMemory = iota
	Redis
	Disk
)

// NewCache is a constructor that returns a Cache interface
func NewCache[K comparable, V any](cacheType int) Cache[K, V] {
	switch cacheType {
	case InMemory:
		return &InMemoryCache[K, V]{data: make(map[K]V)}
	case Redis:
		return &RedisCache[K, V]{data: make(map[K]V)}
	case Disk:
		return &DiskCache[K, V]{data: make(map[K]V)}
	default:
		return nil
	}
}
