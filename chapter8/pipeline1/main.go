package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 0; i < 15; i++ {
			naturals <- i
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// for {
	// 	res, ok := <-squares
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(res)
	// }

	for res := range squares {
		fmt.Println(res)
	}
}
