package main

import (
	"fmt"
	"strconv"
	"sync"
)

// test worker pool implementation
type Work struct {
	//for my example every work will be nothing but just printing a string
	val string
}

type WorkerPool struct {
	jobChannel chan Work
	wg         *sync.WaitGroup
}

// NewWorkerPool Initiates a new worker pool with an executable call back function
func NewWorkerPool(numWorkers int, executeFunc func(val any)) *WorkerPool {
	//create a buffered channel of numWorkers
	myJobChannel := make(chan Work, numWorkers)
	//there needs to be a goRoutine which listens to the jobs and takes some action to it
	//starting all the goRoutines
	var myWaitGroup sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		go func() {
			for {
				select {
				case work := <-myJobChannel:
					executeFunc(work.val)
					myWaitGroup.Done()
				}
			}
		}()
	}

	return &WorkerPool{jobChannel: myJobChannel, wg: &myWaitGroup}
}

// SubmitJob submits the job to the worker pool
func (wp *WorkerPool) SubmitJob(job Work) {
	//pushing the job to the job channel
	fmt.Println("received job ", job.val)
	wp.wg.Add(1)
	wp.jobChannel <- job
	fmt.Println("submitted job to job channel")
}

func (wp *WorkerPool) WaitForAllJobs() {
	// wait group waiting -> do not exit until all jobs are completed
	wp.wg.Wait()
}

// I have a worker pool where I can submit jobs

func executeFunc(val any) {
	switch v := val.(type) {
	case string:
		fmt.Println("The value is a string:", v)
	default:
		fmt.Printf("The value is of type %T and has value %v\n", v, v)
	}

}

func main() {
	wp := NewWorkerPool(5, executeFunc)
	for i := 0; i < 10; i++ {
		wp.SubmitJob(Work{val: strconv.Itoa(i)})
	}
	wp.WaitForAllJobs()
}
