package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

func main() {

	fmt.Println("main --> Call the boring generator which has an internal anonymous goroutine.")
	c := boring("A")

	fmt.Println(color.Ize(color.Red, "main --> Just receive from channel 3 times!"))
	for i := 0; i < 3; i++ {
		// The main goroutine will wait for a value to be sent.
		fmt.Println(color.Ize(color.Green, fmt.Sprintf("main ---> received: %q", <-c)))
	}
	fmt.Println("-----------------------------------------------------------------------------|")
	fmt.Println(color.Ize(color.Blue, "We can have more instances of the boring service:"))
	// Channels as a handle on a service:
	// Our boring function returns a channel that lets us communicate with the
	// boring service it provides. We can have more instances of the service.
	fmt.Println("main -------> Call a new boring service.")
	newBoringService := boring("new boring service")
	fmt.Println(color.Ize(color.Red, "main --> Just receive from channel 4 times!"))
	for i := 0; i < 4; i++ {
		// The main goroutine will wait for a value to be sent.
		fmt.Println(color.Ize(color.Green, fmt.Sprintf("main ---> received: %q", <-newBoringService)))
	}
	fmt.Println(color.Ize(color.Red, "main ---> You're both boring; I'm leaving."))
}

// boring is a generator (in this case it returns receive-only channel of strings)
func boring(msg string) <-chan string {
	c := make(chan string) // 0. Create a channel of string

	fmt.Println("generator --> starting the goroutine")
	go func() { // 1. Run the goroutine from inside the generator.
		for i := 0; ; i++ {
			// 2. Use c channel to send any suitable expression to generator.
			expression := fmt.Sprintf("%s:%d", msg, i)
			// The goroutine waits for a receiver to be ready.
			c <- expression
			fmt.Printf("goroutine --> sent: %q\n", expression)

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	fmt.Println("generator --> returning channel") // See the output
	return c                                       // 3. Return the channel to the caller
}
