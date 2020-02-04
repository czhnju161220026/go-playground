package main

import "fmt"

import "goPlayground/chapter5/links"

import "log"

// 在worklist非空时
// 对每个元素调用f，并将f返回的结果加入到worklist后
// 这是bfs的过程
func bfs(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		nextItem := worklist[0]
		worklist = worklist[1:]
		if !seen[nextItem] {
			seen[nextItem] = true
			neighbours := f(nextItem)
			worklist = append(worklist, neighbours...)
		}
	}
}

func craw(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func main() {
	url := "https://golang.org"
	bfs(craw, []string{url})
}
