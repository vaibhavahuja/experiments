package main

import (
	"fmt"
	"math/rand"
	"time"
)

//PS : goRoutine is not a thread!!!!
//There can be only one thread in a program with 1000's of goroutines
//communication between goRoutines -> channels
//channels communicate and synchronize in a single 

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are boring! I am leaving")

}
