package main

import "fmt"

func sum(a ...int) int {
	res := 0
	for _, x := range a {
		res += x
	}
	return res
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
