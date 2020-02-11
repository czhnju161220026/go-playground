package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for i := 0; i < 20; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
	//	不能都是 go，否则会导致main函数的goroutine结束，程序终止
}
