package main

import (
	"fmt"
	"math/rand"
	"time"
)

var versions []string = []string{"v1.0", "v2.0", "v2.1"}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string
type Search func(query string) Result

func main() {
	rand.Seed(time.Now().UnixNano())
	for _, version := range versions {
		start := time.Now()
		results := Google("golang", version)
		elapsed := time.Since(start)
		fmt.Printf("----------%s----------\n", version)
		fmt.Println(results)
		fmt.Println("elapsed time:", elapsed)
	}
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
	}
	return
}
