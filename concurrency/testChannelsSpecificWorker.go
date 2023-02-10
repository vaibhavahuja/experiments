package main

import (
	"fmt"
	"sync"
	"time"
)

// will be implementing some similar workerPool for my task as well!!!!
func main() {
	start := time.Now()
	const numberOfChannels = 5
	const maxQueueSize = 1
	var chanArray [numberOfChannels]chan int
	wg := new(sync.WaitGroup)
	for i := range chanArray {
		chanArray[i] = make(chan int, maxQueueSize)
	}
	//now I have a chanArray
	//spinning a goRoutine for every Channel array?
	//numberOfChannels goRoutines
	var myArray [5][]int
	for i := range chanArray {
		go func(myChannel chan int, channelNumber int) {
			//fmt.Println("starting a channel which will listen on i : ", channelNumber)
			for {
				select {
				case p := <-myChannel:
					//myValue := time.Duration(channelNumber) * time.Second
					//time.Sleep(myValue)
					fmt.Println(p)
					myArray[channelNumber] = append(myArray[channelNumber], p)
					wg.Done()
					//fmt.Println("received value ", p, "on channel : ", channelNumber)
				}
			}
		}(chanArray[i], i)
	}

	for i := 0; i < 30; i++ {
		val := i % numberOfChannels
		chanArray[val] <- i
		wg.Add(1)
	}
	wg.Wait()
	totalTime := time.Since(start)
	fmt.Println(myArray)
	//time.Sleep(60 * time.Second)
	fmt.Println("total execution ", totalTime)

}
