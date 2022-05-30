package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Generator : function that returns a channel

//Think of it as a service, function returns a channel that lets us communicate with the boring service it provides
//We can have more than one instances of the service

func boring2(msg string) <-chan string { //returns receive-only channel of strings
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

//fan in function or multiplexer
//how it works is it listens to both channels and returns one channel as the return value
//internally two independent goroutines are passed
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

type Message struct {
	str string
	//acts as a signaller waits for one channel
	wait chan bool
}

//refer 20.01

func main() {
	//what if we dont want it to be independent and want it to be associated
	//for i := 0; i < 5; i++ {
	//	msg1 := <-c
	//	fmt.Println(msg1.str)

	//}

	//using fanIn function we decouple it
	c := fanIn(boring2("Joe"), boring2("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	//c := boring2("boring!")
	//c2 := boring2("Ann!")
	//for i := 0; i < 5; i++ {
	//	//because of the synchronizing nature of the channels
	//	//it waits before even going to the next one!!
	//	//one can be fast and other can be slow
	//	//to get around that we can either write a fan-in function or Multiplexer
	//	fmt.Println(<-c)
	//	fmt.Println(<-c2)
	//}
	//fmt.Println("you are boring! I am leaving")
}
