package main

import (
	"fmt"
	"log"
	"os"

	intro "github.com/goplateau/concurrency/oreilly2017/ch01_intro"
)

type Chapter struct {
	id     int
	title  string
	topics []Topic
}

type Topic string

var Chapters = []Chapter{
	{1, "intro", []Topic{"race"}},
	{2, "csp", []Topic{}},
}

func main() {
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case Chapters[0].title:
			intro.Race()
		case Chapters[1].title:
			log.Fatal("not implemented")
		default:
			chapterHelp()
		}
	} else {
		chapterHelp()
	}
}

func chapterHelp() {
	fmt.Printf("Specify a chapter title using\n 	go run . <")
	for i, ch := range Chapters {
		switch i {
		case len(Chapters) - 1:
			fmt.Printf("%s", ch.title)
		default:
			fmt.Printf("%s|", ch.title)
		}
	}
	fmt.Println(">")
}
