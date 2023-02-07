package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	//limiter channel will receive a value every 200 milliseconds
	limiter := time.Tick(1 * time.Second)
	for req := range requests {
		//by blocking on receive from limiter channel we are limiting ourselves
		//to 1 request every 200 milliseconds
		<-limiter
		fmt.Println("request ", req, time.Now())
	}

	//may want to allow short bursts of our requests
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1 * time.Second) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
