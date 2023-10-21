package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

// will create new cache object
// will keep on updating its value whenever I want

func main() {
	c := cache.New(10*time.Second, 10*time.Minute)
	userList := []User{{
		Name: "Vaibhav",
		Age:  25,
	}, {
		Name: "Ahuja",
		Age:  25,
	},
	}
	testObject := &TestObject{UserList: userList}
	//c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("foo2", testObject, 5010*time.Millisecond)

	for i := 0; i < 15; i++ {
		time.Sleep(1 * time.Second)
		//foo, found := c.Get("foo")
		foo2, found2 := c.Get("foo2")
		//if found {
		//	fmt.Println(foo)
		//} else {
		//	fmt.Println("error")
		//}
		if found2 {
			fmt.Println(foo2)
		} else {
			fmt.Println("error")
		}
	}

}
