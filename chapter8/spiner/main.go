package main

import (
	"fmt"
	"time"
)

func spiner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	go spiner(100 * time.Millisecond)
	res := fib(45)
	fmt.Printf("\rFib(45)= %d\n", res)
}
