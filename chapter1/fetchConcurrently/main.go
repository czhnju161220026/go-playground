package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func fetch(url string, channel chan<-string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	seconds := time.Since(start).Seconds()
	channel <- fmt.Sprintf("%.2fs %7d %s", seconds, nbytes, url)
}

func main() {
	urls := [3] string {"http://www.baidu.com", "http://www.apple.cn", "http://gopl.io"}
	ch := make(chan string)
	for _, url := range urls{
		go fetch(url, ch)
	}
	for range urls{
		fmt.Println(<-ch)
	}

}

