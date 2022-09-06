package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

// The Go approach: Don't communicate by sharing memory, share memory by communicating.
func main() {
	c := make(chan string)
	// A channel connects the main and boring goroutines so they can communicate.
	// A sender and receiver must both be ready to play their part in the communication.
	// Otherwise we wait until they are. Thus channels both communicate and synchronize.
	fmt.Println("main --> launch a goroutine with an input channel")
	go boring("A", c)
	/* Go channels can also be created with a buffer (See make function for more details).
	- Buffering removes synchronization.
	- Buffered channels can be important for some problems but they are more subtle to
	reason about.
	*/
	fmt.Println(color.Ize(color.Red, "main --> Just receive from channel 10 times!"))
	for i := 0; i < 10; i++ {
		// The main goroutine will wait for a value to be sent.
		fmt.Println(color.Ize(color.Green, fmt.Sprintf("main-> received: %q", <-c)))
	}
	fmt.Println(color.Ize(color.Red, "main --> You're boring; I'm leaving."))
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// Use c channel to send any suitable expression to main goroutine.
		expression := fmt.Sprintf("%s:%d", msg, i)
		c <- expression
		// The boring goroutine waits for a receiver to be ready.
		fmt.Printf("boring --> sent: %q\n", expression)

		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
