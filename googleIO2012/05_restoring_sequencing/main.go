package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

type Message struct {
	body  string
	block chan bool
}

func main() {
	rand.Seed(time.Now().UnixNano())
	services := []<-chan Message{generator("A"), generator("B"), generator("C")}
	fmt.Printf("main -----------> %d boring services have been made.\n", len(services))

	fmt.Println("main -----------> call fan-in")
	c := fanIn(services)

	fmt.Println("main -----------> Just receive channel value 2 times!")

	for i := 0; i < 2; i++ {
		// The main goroutine will wait for a value to be sent.
		for j := 0; j < len(services); j++ {
			msg := <-c
			fmt.Println(color.Ize(color.Green, fmt.Sprintf("main -----> received: %q", msg.body)))
			<-msg.block // reset channel, stop blocking
		}
	}
	fmt.Println(color.Ize(color.Green, "main ----> sequece of the received channels restored."))
	fmt.Println("You're all boring; I'm leaving.")
}

// We can use a fan-in function to let whosoever is ready to talk.
func fanIn(services []<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := 0; i < len(services); i++ { // launch all goroutines while loops to continuously pipe to new channel
		go func(j int) {
			for {
				c <- <-services[j]
			}
		}(i) // Here i is declared as an explicit parameter of the func literal.
		// It will be taken as a different "j" for each go routine, no matter when
		// the go routine is executed and independently of the current value of i.
	}
	return c
}

func generator(msg string) <-chan Message { // Returns receive-only channel of Message.

	c := make(chan Message)    // 0. Create a channel
	blocked := make(chan bool) // channel within channel to control exec (default: false)

	fmt.Printf("generator %q --> starting the goroutine", msg)
	go func() { // 1. Run the anonymous goroutine from inside the generator.
		for i := 0; ; i++ {
			// 2. Use c channel to send any suitable Message to generator.
			expression := fmt.Sprintf("%s:%d", msg, i)
			// The goroutine waits for a receiver to be ready.
			c <- Message{expression, blocked}
			fmt.Printf("goroutine %q --> sent body: %q\n", msg, expression)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			// Block by waiting for input
			blocked <- true
		}
	}()
	fmt.Printf("generator %q --> returning channel\n", msg) // See the output
	return c                                                // 3. Return the channel to the caller
}
