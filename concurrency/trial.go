package main

import "fmt"

type Numbers struct {
	a int
	b int
}

func addNumbers(num Numbers) <-chan int {
	c := make(chan int)
	go func() {
		c <- num.a + num.b
	}()
	return c
}

func main() {
	numbers := &Numbers{
		a: 12,
		b: 32,
	}
	c := addNumbers(*numbers)
	fmt.Println(<-c)
}
