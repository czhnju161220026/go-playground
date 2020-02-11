package main

import (
	"fmt"
	"goPlayground/chapter5/links"
	"log"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenlinks := make(chan string)

	go func() { worklist <- []string{"https://github.com"} }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenlinks {
				list := crawl(link)
				go func() { worklist <- list }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenlinks <- link
			}
		}
	}
}
