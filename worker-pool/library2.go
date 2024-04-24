package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alitto/pond"
)

func main() {

	// Create a worker pool -> extra tasks will be blocked?
	// I mean can think of it
	pool := pond.New(2, 20)
	//do I really want to stop and wait? -> yes
	// Can always initialise a new worker pool
	defer pool.StopAndWait()
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx, cancel := context.WithCancel(ctx1)

	go func(cancelFunc context.CancelFunc) {
		time.Sleep(5 * time.Second)
		cancelFunc()
		fmt.Println("context is cancelled")
	}(cancel1)

	// Create a task group associated to a context
	group, ctx := pool.GroupContext(ctx)

	var i int
	for i < 50 {
		val := i
		//cancelFunc := cancel
		//fmt.Println(i)
		fmt.Println("submitting job ", val)
		group.Submit(func() error {
			//fmt.Println(val)
			//do I need the response from the task -> not really
			// Will implement this in my task
			err := doSomeWork(val)
			if err != nil {
				cancel()
			}
			return err
		})
		fmt.Println("done w submit job ", val)

		i++
	}

	err := group.Wait()
	//checking if it was all gracefully ended - it was
	if err != nil {
		fmt.Println("Error with some jobs: ", err.Error())
	} else {
		fmt.Println("Successfully fetched all URLs")
	}

	fmt.Println("all tasks are completed now it seems")
}

func doSomeWork(i int) error {
	//simulating every job to take 2 seconds for execution
	time.Sleep(2 * time.Second)
	if i == 10 {
		panic("test painic")
	}
	fmt.Println("job completed -> ", i)

	return nil
}
