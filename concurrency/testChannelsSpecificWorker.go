package main

import (
	"fmt"
	"sync"
	"time"
)

//will be implementing some similar workerPool for my task as well!!!!
func main() {
	const numberOfChannels = 4
	var chanArray [numberOfChannels]chan int
	wg := new(sync.WaitGroup)
	for i := range chanArray {
		chanArray[i] = make(chan int)
	}
	//now I have a chanArray
	//spinning a goRoutine for every Channel array?
	//numberOfChannels goRoutines
	for i := range chanArray {
		go func(myChannel chan int, channelNumber int) {
			fmt.Println("starting a channel which will listen on i : ", i)
			for {
				select {
				case p := <-myChannel:
					//adding delay - to test?
					myValue := time.Duration(numberOfChannels-channelNumber) * time.Second
					time.Sleep(myValue)
					fmt.Println("sleeping for", myValue, "seconds")
					time.Sleep(time.Second)
					fmt.Println("received value ", p, "on channel : ", channelNumber)
					wg.Done()
				}
			}
		}(chanArray[i], i)
	}

	for i := 0; i < 10; i++ {
		val := i % numberOfChannels
		//now what if it already has some value? Let me spin a goRoutine which will do it
		//will try without goRoutine first
		//this should throw an error or will get stuck
		//again spinning infinitegoRoutines is it the right thing?
		chanArray[val] <- i
		wg.Add(1)

	}
	wg.Wait()

}
