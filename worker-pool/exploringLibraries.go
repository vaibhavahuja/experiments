package main

import (
	"fmt"
	"github.com/gammazero/workerpool"
)

var globalWorkerPool *workerpool.WorkerPool

type Work struct {
	downloadPath string
	uploadPath   string
	requestId    string
}

func main() {
	//init global worker pool
	globalWorkerPool = workerpool.New(2)

	pushTasksA(10, "nameA")
	pushTasksA(10, "nameB")
	globalWorkerPool.StopWait()
}

func pushTasksA(numTasks int, name string) {
	var i int
	for i < numTasks {
		task := Work{
			requestId: fmt.Sprintf("%s.%d", name, i),
		}
		globalWorkerPool.Submit(func() {
			fmt.Println("Handling request:", task.requestId)
		})
		i++
	}
}
