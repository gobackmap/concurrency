package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	// Anonymous function as a goroutine
	go func(x int) {
		fmt.Printf("Ananymous: %d\n", x)
	}(10)

	// Function as a goroutine
	go printMe(15)

	// Multiple goroutines
	var (
		count int
		err   error
	)
	if len(os.Args) > 1 {
		count, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Going to create %d goroutines...\n", count)
	}
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("multiple goroutines: %d\n", x)
		}(i)
	}
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}

func printMe(x int) {
	fmt.Printf("Simple Function: %d\n", x)
}
