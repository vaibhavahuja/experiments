package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

//this concept is pretty new and cool it's mainly a function returning a func
//to see more about it is called functionClosures because uses a value outside the function
//idea is if a function is inside a function which uses the value of outside function it's called function closure

type Search func(query string) Result

var (
	Image  = fakeSearch("image")
	Video  = fakeSearch("video")
	Web1   = fakeSearch("web1")
	Web2   = fakeSearch("web2")
	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2")
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
)

//fakeSearch(kind) -> returns a func(queryString)
//to call call it like thisfakeSearch("web")(query)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

//GoogleV1 why do it using blocking calls?
func GoogleV1(query string) (results []Result) {
	results = append(results, fakeSearch("web")(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

//GoogleV2 use goRoutines
//come in order, running concurrently and in order
//waiting only for the slowest -> this is parallel but without mutexes or locks
//Again some queries can take lots of time, do we really have that much time? Solved in GoogleV3
func GoogleV2(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- fakeSearch("web")(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()
	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

//What if we do not want to wait for the complete duration?
//GoogleV3 -> solves the problem mentioned above

func GoogleV3(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			//times out the responses which are not received
			//max upto 80ms!
			fmt.Println("timed out mate")
			return
		}
	}
	return
}

//How do we avoid discarding results from slow servers?
//Replicate the servers. send request to multiple replicas and use the first response
//Launch multiple searches at same time and return the first one which returns the response

// How do we avoid discarding result from the slow server.
// We duplicate to many instance, and perform parallel request.

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	for i := range replicas {
		go func(idx int) {
			c <- replicas[idx](query)
		}(i)
	}
	// the magic is here. First function always waits for 1 time after receiving the result
	return <-c
}

//Just a toy example
//See how we are using goRoutines to build parallel robust thing
//None of them use threads, no callBacks (nodeJs)
//individual elements are straightforward individual code
//very cool
func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	//results := GoogleV1("golang")
	//results := GoogleV2("golang")

	//having replicas below
	//result := First("golang", fakeSearch("replica 1"), fakeSearch("replica 2"), fakeSearch("replica 1"), fakeSearch("replica 1"))
	results := GoogleV3("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	//fmt.Println(result)
	fmt.Println(elapsed)
}
