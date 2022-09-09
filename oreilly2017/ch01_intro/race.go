package intro

import (
	"fmt"

	"github.com/fatih/color"
)

func Race() {
	const N int = 50
	var data int
	fmt.Println("==============================================")
	fmt.Println(" Race Condition Exmample")
	fmt.Printf(" Number of Iteration=%d, shared data=%d\n", N, data)
	fmt.Println(" Three possible outcomes:")
	fmt.Println("   -Nothing to print")
	fmt.Println("   -data=0")
	fmt.Println("   -data=1")
	fmt.Println(" Just run the program again and again!")
	fmt.Println("==============================================")

	fgGreen := color.New(color.FgGreen)
	fgBlue := color.New(color.FgBlue, color.Underline)
	go func() {
		fgGreen.Println("\ngoroutine >>> the value of data is going to increase...")
		data++
		fgGreen.Println("goroutine >>> the value of data increased...")
	}()
	for i := 0; i < N; i++ {
		fmt.Printf(" iteration %d... ", i)
		if data == 0 {
			fgBlue.Printf("| it=%d, data=%v |", i, data)
		}
	}
	fmt.Println()
}
