package main

import (
	"fmt"
	"os"

	intro "github.com/gobeen/concurrency/oreilly2017/ch01_intro"
)

type Chapter struct {
	id     int
	title  string
	topics []Topic
}

type Topic struct {
	title       string
	description string
	callback    func()
}

var Chapters = []Chapter{
	{1, "intro", []Topic{
		{"race", "race condition", intro.Race},
		{"sync", "memory access synchronization", intro.MemoryAccessSynchronization},
		{"deadlock", "deadlock", intro.Deadlock},
		{"livelock", "livelock", intro.Livelock},
	}},
	{2, "csp", []Topic{}},
}

var args []string

func main() {
	args = os.Args
	if len(args) > 1 {
		for _, chapter := range Chapters {
			if chapter.title == args[1] {
				checkAndRunTopic(chapter)
				break
			}
		}
		fmt.Println("no such chapter!")
		helpChapter()
	} else {
		helpChapter()
	}
}

func checkAndRunTopic(chapter Chapter) {
	if len(args) > 2 {
		for _, topic := range chapter.topics {
			if topic.title == args[2] {
				topic.callback()
				break
			}
		}
		fmt.Printf("no such topic in %q chapter.\n", chapter.title)
		helpTopic(chapter)
	} else {
		helpTopic(chapter)
	}
}

func helpChapter() {
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
	os.Exit(0)
}

func helpTopic(chapter Chapter) {
	fmt.Printf("Specify a topic using\n 	go run . %s <", chapter.title)
	for i, topic := range chapter.topics {
		switch i {
		case len(chapter.topics) - 1:
			fmt.Printf("%s", topic.title)
		default:
			fmt.Printf("%s|", topic.title)
		}
	}
	fmt.Println(">")
	os.Exit(0)
}
