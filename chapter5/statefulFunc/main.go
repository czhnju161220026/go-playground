package main

import "fmt"

func squres() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squres()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
