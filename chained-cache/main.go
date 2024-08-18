package main

import "fmt"

func main() {
	//initialising 3 caches here
	cache1 := NewCache[string, int64](0)
	cache2 := NewCache[string, int64](1)
	cache3 := NewCache[string, int64](2)

	//chainedCache := NewChainedCacheFirst[string, int64](cache1, cache2, cache3)
	//
	//cache1.AddKey("hey", 4)
	//cache2.AddKey("hey2", 5)
	//cache3.AddKey("hey", 6)

	chainedCache := NewChainedCacheSecond[string, int64](cache1, cache2, cache3)

	err := cache2.AddKey("hey", 1)
	if err != nil {
		fmt.Errorf("error while adding key")
	}

	val, err := chainedCache.Get("hey", 0)
	if err != nil {
		fmt.Println("did not find value in chained cache - ", err.Error())
		return
	}
	fmt.Println(" value is ", val)

	//now check if its there in all 3 caches?

	for i, cache := range chainedCache.layers {
		val, err = cache.GetVal("heyggh")
		fmt.Printf("found in cache - %d value - %d", i, val)
		fmt.Println()
	}

}
