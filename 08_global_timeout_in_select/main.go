package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

func main() {
	log.Println(color.Ize(color.Red, "main --> goroutine timeout: 5 second"))
	// Set a timeout for the entire conversation (not for each message)
	timeout := time.After(5 * time.Second)

	c := generator("A")
	for {
		select {
		case s := <-c:
			fmt.Println(color.Ize(color.Green, fmt.Sprintf("main --> received: %q", s)))
		case <-timeout:
			log.Println(color.Ize(color.Red, "main --> too long conversation: timeout exeeded!"))
			return
		}
	}
}

func generator(msg string) <-chan string { // returns receive-only channel
	ch := make(chan string)
	go func() { // anonymous goroutine
		for i := 0; ; i++ {
			expression := fmt.Sprintf("%s:%d", msg, i)
			ch <- expression
			fmt.Printf("goroutine %q --> sent: %q\n", msg, expression)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}
