// implementing a worker pool using goroutines and channels
package main

import (
	"fmt"
	"log"
	"time"
)

//now all odd numbers should process one by one and all even should be processed in order
func worker4(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job", j)
		//sleeping a second to simulate an expensive job
		fmt.Println("worker ", id, "working on ", j)
		time.Sleep(time.Second)
		//fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func collectResults(results <-chan int){
	<-results
}

// diving the jobs amongst different workers for concurrency
//even will always be in sync and odd will always be in sync
//causal ordering I guess
func main() {
	const numJobs = 10
	// is it good to have buffered channel list?
	var jobArray = []chan int {
		make(chan int, 15),
		make(chan int, 15),
	}
	results := make(chan int)

	//jobs := make(chan int, numJobs)
	//starting up 3 workers each consuming from a different channel it seems
	for w := 1; w <= 2; w++ {
		go worker4(w, jobArray[w-1], results)
	}

	//sending 10 jobs
	for j := 1; j <= numJobs*2; j++ {
		log.Println(j)
		if j%2 == 0 {
			jobArray[0] <- j
		} else {
			jobArray[1] <- j
		}
	}

	//avoiding using buffered channels here
	for j := 1; j <= 10; j++ {
		go collectResults(results)
	}

	time.Sleep(20*time.Second)
	//close(jobArray)

	//collect all results
	//alternative way is to use waitGroup ?
	//time.Sleep(10*time.Second)
	//for a := 1; a <= numJobs; a++ {
	//	<-results
	//}
}
