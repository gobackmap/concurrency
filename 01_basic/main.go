package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sleep := 2 * time.Second
	log.Println(color.Ize(color.Red, fmt.Sprintf("main --> waiting for %s", sleep.String())))
	go boring("A")
	time.Sleep(sleep)
	log.Println(color.Ize(color.Red, "main --> You're boring; I'm leaving."))
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("boring goroutine -> %s:%d\n", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
