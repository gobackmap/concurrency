package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	timeout := 1 * time.Second
	fmt.Println(color.Ize(color.Red, fmt.Sprintf("main --> goroutine timeout: %s", timeout.String())))
	c := generator("A")
	for {
		select {
		case s := <-c:
			fmt.Println(color.Ize(color.Green, fmt.Sprintf("main --> received: %q", s)))
		case <-time.After(1 * time.Second):
			fmt.Println(color.Ize(color.Red, "main --> too late, timeout exeeded!"))
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
			time.Sleep(time.Duration(rand.Intn(1.5e3)) * time.Millisecond)
		}
	}()
	return ch
}
