package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Launched")
}

func listenKeyboard(abort chan<- struct{}) {
	os.Stdin.Read(make([]byte, 1))
	abort <- struct{}{}
}

func main() {
	abort := make(chan struct{})
	go listenKeyboard(abort)
	tick := time.NewTicker(1 * time.Second)
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		//多路复用
		//如果同时满足多个情况，会随机选择
		select {
		case <-tick.C:
			//do nothing
		case <-abort:
			fmt.Println("Abort!")
			tick.Stop()
			return
		}
	}
	launch()
}
