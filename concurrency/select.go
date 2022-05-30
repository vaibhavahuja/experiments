package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Using select statement with concurrency
//works like a switch
// another way to handle multiple channels
// evaluates all the channels
// blocks until one of them is available to run
// if no default, blocks until any one of them is available!!!!

//what if multiple channels are available? select chooses any one at random!!!!

//func fanIn(input1, input2 <-chan string) <-chan string {
//	c := make(chan string)
//	go func() {
//		for {
//			c <- <-input1
//		}
//	}()
//	go func() {
//		for {
//			c <- <-input2
//		}
//	}()
//
//	return c
//}

//Rewriting above fanIn function with select statement

func fanInSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

//you can use to select for timeOut

func boring3(msg string) <-chan string { //returns receive-only channel of strings
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
func main() {
	////using fanIn function we decouple it
	//c := fanInSelect(boring3("Joe"), boring3("Ann"))
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-c)
	//}

	//using timeOuts using select
	c := boring3("Joe")
	timeout := time.After(3 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("you talk too much")
			return

		}
	}

	//Another thing you can do with select is that
	//Ask Joe to stop
	//have a quit channel but before quitting actually do a cleanUp!!
}
