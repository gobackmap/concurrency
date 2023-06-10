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
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	fmt.Printf("Waiting for %d number of goroutines to be done...\nNumber of goroutines: %d\n", count, runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("\nDone...only %d goroutine remained: main goroutine!\n", runtime.NumGoroutine())
	fmt.Println("Exit...")
}
