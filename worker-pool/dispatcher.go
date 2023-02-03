package main

import (
	"fmt"
)

type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)
	return Dispatcher{
		workerPool: workerPool,
		maxWorkers: maxWorkers,
		jobQueue:   jobQueue,
	}
}

func (d *Dispatcher) run() {
	for i := 0; i < d.maxWorkers; i++ {
		//creating new worker goRoutines
		fmt.Printf("starting worker %d \n", i+1)
		worker := NewWorker(d.workerPool, i+1)
		worker.start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			//it keeps spinning a new goRoutine
			//doesn't that mean infinite goRoutines again to what they were doing earlier?
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				workerJobQueue := <-d.workerPool
				//adding job to workerJobQueue
				workerJobQueue <- job
			}(job)
		}
	}
}
