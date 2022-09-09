package main

import (
	"fmt"
	"math/rand"
	"time"
)

var versions = []string{"v1.0", "v2.0", "v2.1", "v3.0"}

var (
	Web    = fakeSearch("web")
	Web2   = fakeSearch("web")
	Image  = fakeSearch("image")
	Image2 = fakeSearch("image")
	Video  = fakeSearch("video")
	Video2 = fakeSearch("video")
)

type Result string
type Search func(query string) Result

func main() {
	rand.Seed(time.Now().UnixNano())
	for _, version := range versions {
		start := time.Now()
		results := Google("golang", version)
		elapsed := time.Since(start)
		fmt.Printf("--------------------%s--------------------\n", version)
		fmt.Println(results)
		fmt.Println("elapsed time:", elapsed)
	}
	// Using the first function
	start := time.Now()
	results := First("golang", fakeSearch("replica 1"), fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Println("----------Using The First Function----------")
	fmt.Println(results)
	fmt.Println("elapsed time:", elapsed)
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string, version string) (results []Result) {
	switch version {
	case "v1.0":
		// Invoke Web, Image, and Video searches serially, appending them to the results slice.
		results = append(results, Web(query))
		results = append(results, Image(query))
		results = append(results, Video(query))
	case "v2.0":
		// Run the Web, Image, and Video searches concurrently, and wait for all results
		// No locks. No condition variables. No callbacks.
		c := make(chan Result)
		go func() { c <- Web(query) }()
		go func() { c <- Image(query) }()
		go func() { c <- Video(query) }()

		for i := 0; i < 3; i++ {
			result := <-c
			results = append(results, result)
		}
	case "v2.1":
		// Don't wait for slow servers. No locks. No condition variables. No callbacks.
		c := make(chan Result)
		go func() { c <- Web(query) }()
		go func() { c <- Image(query) }()
		go func() { c <- Video(query) }()

		timeout := time.After(80 * time.Millisecond)
		for i := 0; i < 3; i++ {
			select {
			case result := <-c:
				results = append(results, result)
			case <-timeout:
				fmt.Println("timed out")
				return
			}
		}
	case "v3.0":
		// Reduce tail latency using replicated search servers.
		c := make(chan Result)
		go func() { c <- First(query, Web, Web2) }()
		go func() { c <- First(query, Image, Image2) }()
		go func() { c <- First(query, Video, Video2) }()

		timeout := time.After(80 * time.Millisecond)

		for i := 0; i < 3; i++ {
			select {
			case result := <-c:
				results = append(results, result)
			case <-timeout:
				fmt.Println("timed out")
				return
			}
		}
	}
	return
}

// First replicates the servers to avoid discarding results from slow ones.
// It sends requests to multiple replicas, and uses the first response.
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}
