package main

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) Result

type Result string

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Run the Web, Image, and Video searches concurrently, and wait for all results.
// No locks. No condition variables. No callbacks.
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()
	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}


// https://golang.org/pkg/runtime/pprof/
func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	var results []Result
	results = Google("golang")
	fmt.Println(results)
	
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
