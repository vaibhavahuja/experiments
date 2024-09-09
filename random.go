package main

import (
	"fmt"
	"time"
)

func print(i, max int, p chan bool) {
	//block until some value is received
	for q := 0; q < max; q++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(i)
	}
	p <- true
}

// now i want in order
func main() {
	//var i int
	p := []int{1, 2, 3, 4, 5, 6, 7, 8}
	now := time.Now()
	var ans int
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			ans += p[i] * p[j]
		}
	}
	fmt.Println(ans)

	since := time.Since(now)
	fmt.Println("total time is ", since)
}
