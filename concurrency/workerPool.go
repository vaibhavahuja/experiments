// implementing a worker pool using goroutines and channels
package main

import (
	"fmt"
	"time"
)

func worker9(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job", j)
		fmt.Println("worker ", id, "working on", j)

		//sleeping a second to simulate an expensive job
		//time.Sleep(100 * time.Millisecond)
		time.Sleep(2 * time.Second)
		//fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

// diving the jobs amongst different workers for concurrency
func main() {
	const numJobs = 5
	jobs := make(chan int, 2)
	results := make(chan int)

	//starting up 3 workers
	for w := 1; w <= 3; w++ {
		go worker9(w, jobs, results)
	}

	go func() {
		for a := 1; a <= 25; a++ {
			//fmt.Println("length of results is ", len(results))
			<-results
		}
	}()
	//sending 20 jobs -> 3 workers are pooled up
	for j := 1; j <= 10; j++ {
		fmt.Println(" i am on j ", j)
		select {
		case jobs <- j:
			//put j in jobs unless it is full
		default:
			fmt.Println("skipping ", j, "length of channel is ", len(jobs))
			//discard the value
			time.Sleep(500 * time.Millisecond)
		}
		//jobs <- j
	}
	close(jobs)

	close(results)
}
