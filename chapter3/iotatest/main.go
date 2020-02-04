package main

import "fmt"

const (
	_ = 1 << (iota * 10)
	kiB
	miB
	giB
	tiB
	piB
)

func f(i int) int {
	res := 1
	for index := 0; index < i; index++ {
		res *= 1000
	}
	return res
}

func main() {
	fmt.Printf("kib = %d \nmiB = %d \ngib = %d \ntiB = %d\npiB = %d\n", kiB, miB, giB, tiB, piB)
}
