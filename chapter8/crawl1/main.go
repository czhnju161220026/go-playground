package main

import (
	"fmt"
	"goPlayground/chapter5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

// 无限制的并行导致创建了太多网络连接，程序最终因为资源耗尽退出
func main() {
	worklist := make(chan []string)

	//从命令行参数开始
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
