package main

import (
	"errors"
	"fmt"
)

// InMemoryCache implementation
type InMemoryCache[K comparable, V any] struct {
	data map[K]V
}

func (c *InMemoryCache[K, V]) AddKey(key K, value V) error {
	fmt.Println("adding key to inMemory Cache")
	c.data[key] = value
	return nil
}

func (c *InMemoryCache[K, V]) GetVal(key K) (V, error) {
	fmt.Println("fetching val from inMemory cache")
	val, exists := c.data[key]
	if !exists {
		var zero V
		return zero, errors.New("key not found")
	}
	return val, nil
}

// RedisCache implementation
type RedisCache[K comparable, V any] struct {
	// Simulated redis storage, in real-world it would be different
	data map[K]V
}

func (r *RedisCache[K, V]) AddKey(key K, value V) error {
	fmt.Println("adding key to redis Cache")
	r.data[key] = value
	return nil
}

func (r *RedisCache[K, V]) GetVal(key K) (V, error) {
	fmt.Println("fetching val from redis cache")
	val, exists := r.data[key]
	if !exists {
		var zero V
		return zero, errors.New("key not found")
	}
	return val, nil
}

// DiskCache implementation
type DiskCache[K comparable, V any] struct {
	// Simulated disk storage
	data map[K]V
}

func (d *DiskCache[K, V]) AddKey(key K, value V) error {
	fmt.Println("adding key to disk Cache")
	d.data[key] = value
	return nil
}

func (d *DiskCache[K, V]) GetVal(key K) (V, error) {
	fmt.Println("fetching val from disk cache")
	val, exists := d.data[key]
	if !exists {
		var zero V
		return zero, errors.New("key not found")
	}
	return val, nil
}
