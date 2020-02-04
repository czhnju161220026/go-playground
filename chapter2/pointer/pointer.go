package main

import (
	"fmt"
)

func main() {
	var x int
	px := &x
	fmt.Println(px)
	*px = 2
	fmt.Println(x)
}
