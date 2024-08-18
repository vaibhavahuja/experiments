package main

import (
	"errors"
	"fmt"
)

type ChainedCacheFirst[K comparable, V any] struct {
	//list of caches in the order I want to execute them
	layers []Cache[K, V]
}

func NewChainedCacheFirst[K comparable, V any](cachesInOrder ...Cache[K, V]) *ChainedCacheFirst[K, V] {
	return &ChainedCacheFirst[K, V]{
		layers: cachesInOrder,
	}
}

// GetVal basic implementation - just a simple for loop and checking here
// easy to implement because all caches were same interface
func (chainFirst *ChainedCacheFirst[K, V]) GetVal(key K) (V, error) {
	for _, eachCache := range chainFirst.layers {
		if val, err := eachCache.GetVal(key); err != nil {
			//returns the first value I find in layers of caches
			fmt.Println("error ", err.Error())
		} else {
			fmt.Println("found in chained cache ", val)
			return val, nil
		}
	}

	var zeroVal V
	return zeroVal, errors.New("did not find value at all")
}

func (chainFirst *ChainedCacheFirst[K, V]) AddKey(key K, val V) error {
	//adds the key to all the layers
	for _, cache := range chainFirst.layers {
		if err := cache.AddKey(key, val); err != nil {
			return err
		}
	}
	return nil
}
