package main

import (
	"errors"
	"fmt"
)

//now i will do the recursion way

type ChainedCacheSecond[K comparable, V any] struct {
	//list of caches in the order I want to execute them
	layers []Cache[K, V]
}

func NewChainedCacheSecond[K comparable, V any](cachesInOrder ...Cache[K, V]) *ChainedCacheSecond[K, V] {
	return &ChainedCacheSecond[K, V]{
		layers: cachesInOrder,
	}
}

// Get recursive method which calls next cache and if found
// falls back and updates the existing caches as well
func (chainSecond *ChainedCacheSecond[K, V]) Get(key K, index int) (V, error) {
	currCache := chainSecond.layers[index]

	val, err := currCache.GetVal(key)
	if err != nil {
		//did not find in currCache so will try to find in parent cache
		if index+1 < len(chainSecond.layers) {
			val, err = chainSecond.Get(key, index+1)
			if err == nil {
				//backtracking and updating the parent caches as well
				currCache.AddKey(key, val)
				//successfully found
				fmt.Println("found key in cache ", index+1)
				return val, nil
			} else {
				//found some error
				fmt.Println("key not found")
			}
		}
		//did not find the value after exploring all caches
		var v V
		return v, errors.New("did not find the key in any cache")
	}

	return val, err
}
