package intro

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	fgGreen = color.New(color.FgGreen)
	fgBlue  = color.New(color.FgBlue, color.Underline)
)

func Race() {
	numReadingTry := 50
	var sharedData int

	fmt.Println("==============================================")
	fmt.Println(" Race Condition Exmample")
	fmt.Printf(" Number of trying to read=%d, shared data=%d\n", numReadingTry, sharedData)
	fmt.Println(" Three possible outcomes:")
	fmt.Println("   -Nothing to print")
	fmt.Println("   -data=0")
	fmt.Println("   -data=1")
	fmt.Println(" Just run the program again and again!")
	fmt.Println("==============================================")

	go func() {
		fgGreen.Println("\ngoroutine >>> the value of data is going to increase...")
		sharedData++
		fgGreen.Println("goroutine >>> the value of data increased...")
	}()
	for i := 0; i < numReadingTry; i++ {
		fmt.Printf(" iteration %d... ", i)
		if sharedData == 0 {
			fgBlue.Printf("| it=%d, data=%v |", i, sharedData)
		}
	}
	fmt.Println()
}
