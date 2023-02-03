package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		maxWorkers   = 1
		maxQueueSize = 0
	)

	jobQueue := make(chan Job, maxQueueSize)

	//create and start the dispatcher
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.run()
	//simulating some requests
	jobs := []Job{
		{
			Name:  "first",
			Delay: 3,
		},
		{
			Name:  "second",
			Delay: 3,
		}, {
			Name:  "third",
			Delay: 3,
		},
		{
			Name:  "fourth",
			Delay: 3,
		},
	}

	for _, i := range jobs {
		fmt.Printf("sending job %s to jobQueue at time %v \n", i.Name, time.Now().Second())
		jobQueue <- i
		fmt.Printf("length of jobQueue is %d \n", len(jobQueue))
	}

	time.Sleep(30 * time.Second)
}
