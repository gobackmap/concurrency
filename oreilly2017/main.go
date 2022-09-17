package main

import (
	"fmt"
	"log"
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
		switch args[1] {
		case Chapters[0].title:
			checkTopic(Chapters[0])
		case Chapters[1].title:
			log.Fatal("not implemented")
		default:
			chapterHelp()
		}
	} else {
		chapterHelp()
	}
}

func checkTopic(chapter Chapter) {
	if len(args) > 2 {
		for _, topic := range chapter.topics {
			if topic.title == args[2] {
				topic.callback()
			}
		}
	} else {
		topicHelp(chapter)
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

func topicHelp(chapter Chapter) {
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
}
