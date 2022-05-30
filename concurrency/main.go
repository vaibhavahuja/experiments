// Understanding Channels and how blocking works

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "every 500 ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "every 2 seconds"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	//the below will slow down since it blocks and waits until the channel returns something
	//for {
	//	fmt.Println(<-c1)
	//	fmt.Println(<-c2)
	//}
	// Instead using select case

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func getDataFromChanel(c chan string) {
	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}
}

//writing stuff to a channel
func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
