package main

import (
	"fmt"
	"sync"
	"time"
)

func testBlocking(test <-chan int, quitChan <-chan bool, wg *sync.WaitGroup) {
	for {
		//test is empty as of now!!
		fmt.Printf("printing it time \n")
		select {
		//tested - this blocks teh goRoutine at this place waiting for test to return something
		case p := <-test:
			fmt.Printf("finally received something which is %d \n", p)
		case <-quitChan:
			fmt.Printf("ahh have to quit now \n")
		}
		wg.Done()
	}
}

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(1)
	myChannel := make(chan int)
	quitChan := make(chan bool)
	go testBlocking(myChannel, quitChan, wg)

	time.Sleep(2 * time.Second)
	myChannel <- 7
	wg.Wait()
}
