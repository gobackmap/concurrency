package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	services := []<-chan string{boring("A"), boring("B"), boring("C")}
	fmt.Printf("main --> %d boring services have been made.\n", len(services))

	fmt.Println("main --> call fan-in")
	c := fanIn(services)

	fmt.Println("main --> Just receive channel value 6 times!")
	for i := 0; i < 6; i++ {
		// The main goroutine will wait for a value to be sent.
		fmt.Println(color.Ize(color.Green, fmt.Sprintf("main -----> received: %q", <-c)))
	}
	fmt.Println(color.Ize(color.Blue, "main ----> Notice to sequece of the received channels."))
	fmt.Println("main ----> You're all boring; I'm leaving.")
}

// We can use a fan-in function to let whosoever is ready to talk.
func fanIn(services []<-chan string) <-chan string {
	c := make(chan string)
	for i := 0; i < len(services); i++ {
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
func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string) // 0. Create a channel of string

	fmt.Printf("generator %q-> starting the goroutine\n", msg)
	go func() { // 1. Run the goroutine from inside the generator.
		for i := 0; ; i++ {
			// 2. Use c channel to send any suitable expression to generator.
			expression := fmt.Sprintf("%s:%d", msg, i)
			// The goroutine waits for a receiver to be ready.
			c <- expression
			fmt.Printf("goroutine %q-> sent: %q\n", msg, expression)

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	fmt.Printf("generator %q-> returning channel\n", msg) // See the output
	return c                                              // 3. Return the channel to the caller
}
