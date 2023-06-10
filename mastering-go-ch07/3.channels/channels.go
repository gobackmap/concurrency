package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int, 1)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(ch chan int) {
		defer waitGroup.Done()
		writeToChannel(ch, 10)
		fmt.Println("Exit.")
	}(c)
	for i := 0; i <= 2; i++ {
		readChannel(c)
	}
	waitGroup.Wait()
}

func writeToChannel(ch chan int, val int) {
	ch <- val
	close(ch)
}

func readChannel(ch chan int) {
	value, open := <-ch
	if open {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
	fmt.Println("Read:", value)
}
