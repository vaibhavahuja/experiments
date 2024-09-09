package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch(("image"))
	Video = fakeSearch(("video"))
)

func main() {

}

type Result struct {
}

type Search func(query string) string

func fakeSearch(kind string) Search {
	return func(query string) string {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return fmt.Sprintf("%s result for %q", kind, query)
	}
}
