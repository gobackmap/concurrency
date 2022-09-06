package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)
	ch := generator("A", quit)
	fmt.Println("main --> Throwing a twenty-sided dice!")

	dice := rand.Intn(20)
	fmt.Println(color.Ize(color.Red, fmt.Sprintf("main --> You have %d chances in this channel!", dice)))
	for i := dice; i > 0; i-- {
		fmt.Println(color.Ize(color.Green, fmt.Sprintf("main --> received: %q", <-ch)))
	}
	quit <- true
	fmt.Println(color.Ize(color.Red, "I'm quit!"))
}

func generator(msg string, quit <-chan bool) <-chan string { // returns receive-only channel
	ch := make(chan string)
	go func() { // anonymous goroutine
		for i := 0; ; i++ {
			expression := fmt.Sprintf("%s:%d", msg, i)
			select {
			case ch <- expression:
				// Just sent the expression to the output channel
				fmt.Printf("goroutine %q: send %q\n", msg, expression)
			case <-quit:
				// return if quit channel receive true
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}
