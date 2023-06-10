package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var (
	dir      = "./mastering-go-ch07/2.waitgroup/tmp"
	filename string
)

var (
	start, end int
	err        error
	wg         sync.WaitGroup
)

func main() {
	for i := start; i <= end; i++ {
		wg.Add(1)
		name := fmt.Sprintf("%d.%s", i, filename)
		go func(fName string) {
			defer wg.Done()
			if wErr := os.WriteFile(dir+"/"+fName, []byte("This is a test file"), 0666); wErr != nil {
				log.Println("failed to write file:", dir+"/"+fName)
			}
		}(name)
	}
	wg.Wait()
	fmt.Println("Exit...")
}

func init() {
	if len(os.Args) != 4 {
		log.Fatal("Usage: go run .../waitgroup.go startInt endInt filename")
	}
	for i, v := range os.Args {
		switch i {
		case 1:
			start, err = strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			end, err = strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
		case 3:
			filename = v
		}
	}
	if err = os.Mkdir(dir, 0750); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	fmt.Printf("Going to create %d goroutines for each file creation...\n", start-end+1)

}
