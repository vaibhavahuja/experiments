package main

import (
	"fmt"
	"sync"
)

// class that has just a single instance
// cannot create copies of it. where is it useful? say accessing a db
//or common resource
// why not to create a new global var? Anyone can overwrite it
// causing program to crash.

//singleton pattern is used in such cases; call will return the same object
// no matter how many times you call it
//used say for e.g : same db object shared between different parts of program
// stricter control over global variables is needed

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

// single instance is declared globally but is private, and can only be accessed via this method
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance already created")
		}
	} else {
		fmt.Println("single instance already created")
	}
	return singleInstance
}

// there are also other methods of creating a singleton instance in Go
// init method is only called once per file in a package, so we can be sure
// that only a single instance will be created
//func init() {
//	singleInstance = &single{}
//}

// can also use sync.Once

var once sync.Once

func getInstanceOnce() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("creating single instance now")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("single instance already created.")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 10; i++ {
		//go getInstance()
		go getInstanceOnce()
	}
	fmt.Scanln()
}
