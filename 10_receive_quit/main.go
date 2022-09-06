package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan string)
	ch := generator("A", quit)

	fmt.Println("main --> Throwing a twenty-sided dice!")
	dice := rand.Intn(20)
	fmt.Println(color.Ize(color.Red, fmt.Sprintf("main --> You have %d chances in this channel!", dice)))

	for i := dice; i > 0; i-- {
		fmt.Println(color.Ize(color.Green, fmt.Sprintf("main->received: %q", <-ch)))
	}
	quit <- "Bye!"
	fmt.Printf("main -------> received %q on quit channel.\n", <-quit)
	fmt.Println(color.Ize(color.Red, "main -------> I'm quit!"))
}

func generator(msg string, quit chan string) <-chan string { // returns receive-only channel
	ch := make(chan string)
	go func() { // anonymous goroutine
		for i := 0; ; i++ {
			expression := fmt.Sprintf("%s:%d", msg, i)
			select {
			case ch <- expression:
				// Just sent the expression to the output channel
				fmt.Printf("goroutine--%q: send %q\n", msg, expression)
			case q := <-quit:
				// Say something if quit channel received a string
				fmt.Printf("generator --> received %q on quit channel. \n", q)
				quit <- "Good Luck!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}
