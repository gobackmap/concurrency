package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("<<<<<<<<< Runtime >>>>>>>>>")
	fmt.Printf(" Compiler: %s\n Machine: %s\n Version: %s\n GOMAXPROCS:%d\n",
		runtime.Compiler, runtime.GOARCH, runtime.Version(), runtime.GOMAXPROCS(0),
	)
}
