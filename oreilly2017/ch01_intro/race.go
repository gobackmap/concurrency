package main

import (
	"fmt"
	"time"
)

func main() {
	const N int = 50
	var data int
	fmt.Println("==============================================")
	fmt.Println(" Race Condition Exmample")
	fmt.Printf(" Number of Iteration=%d, shared data=%d\n", N, data)
	fmt.Println(" Just run the program again and agian!")
	fmt.Println("==============================================")

	go func() {
		fmt.Println("goroutine >>> the value of data is going to increase...")
		data++
		fmt.Println("goroutine >>> the value of data increased...")
	}()
	for i := 0; i < N; i++ {
		if data == 0 {
			fmt.Printf("it=%d, data=%v\n", i, data)
		}
	}
	time.Sleep(time.Second)
}
