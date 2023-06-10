package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	// Multiple goroutines
	var (
		count int
		err   error
		wg    sync.WaitGroup
	)
	if len(os.Args) > 1 {
		count, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Going to create %d goroutines...\n", count)
	}
	if count == 0 {
		log.Println("at least one positive integer arg is required!")
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(x int) {
			fmt.Printf("%d ", x)
			wg.Done()
		}(i)
	}
	fmt.Printf("Waiting for %d number of goroutines to be done...\n", runtime.NumGoroutine()-1)
	wg.Wait()
	fmt.Println("Exit...")
}
