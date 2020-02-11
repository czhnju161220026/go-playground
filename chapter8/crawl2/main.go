package main

import (
	"fmt"
	"goPlayground/chapter5/links"
	"log"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} //获取令牌
	list, err := links.Extract(url)
	<-tokens //释放令牌
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int
	go func() { worklist <- []string{"https://github.com"} }()
	n++

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
