package main

import "fmt"

// takes two int channels, increases the right channel by 1, and stores it into left channel
func chain(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go chain(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
