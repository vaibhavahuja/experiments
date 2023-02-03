package main

import (
	"fmt"
	"time"
)

type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
}

// NewWorker creates a new worker inside the existing workerPool
func NewWorker(workerPool chan chan Job, id int) Worker {
	return Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}

// starts a new worker and works on the jobs given to it
func (w Worker) start() {
	go func() {
		for {
			//adding job queue to worker pool - why is this being done?
			//understood why it is being done

			// in the dispatcher it fetches the jobQueue from workerPool
			//and then assigns the job to that jobQueue
			//should it be if jobQueue is empty only then send?
			//let me try

			//if len(w.jobQueue) == 0 {
			//oh I get it, it will always be empty here
			//fmt.Printf("length of the job queue is %d \n", len(w.jobQueue))
			//registering worker to workerPool
			w.workerPool <- w.jobQueue
			//}

			select {
			//does it block? Yes it does, wait here for a job!!
			case job := <-w.jobQueue:
				//dispatcher has added a new job so work on that job
				fmt.Printf("worker %d started %s, blocking for %d seconds \n", w.id, job.Name, job.Delay)
				time.Sleep(job.Delay * time.Second)
				fmt.Printf("worker %d completed %s! \n", w.id, job.Name)
			case <-w.quitChan:
				fmt.Printf("worker %d stopping", w.id)
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}
