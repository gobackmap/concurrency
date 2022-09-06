package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("main ---> call fan-in with two A and B generators")
	ch := fanIn(generator("A"), generator("B"))

	fmt.Println("main ---> Just receive channel value 10 times!")
	for i := 0; i < 10; i++ {
		fmt.Println(color.Ize(color.Green, fmt.Sprintf("main ---> received: %q", <-ch)))
	}
	fmt.Println(color.Ize(color.Blue, "main ----> Notice to sequece of the received channels."))

}

// fanIn receives two read-only channels, returns receive-only one (it is a generator)
func fanIn(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				ch <- s
			case s := <-ch2:
				ch <- s
			}
			/* 	https://go.dev/ref/spec#Select_statements
			If one or more of the communications [expressed in the case statements]	can
			proceed, a single one that can proceed is chosen via a uniform pseudo-random
			selection. Otherwise, if there is a default case, that case is chosen. If there
			is no default case, the "select" statement blocks until at least one of
			the communications can proceed.
			*/
		}
	}()
	return ch
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
