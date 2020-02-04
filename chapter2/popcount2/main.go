package main

import (
	"fmt"
)

func popCountV2(x uint64) int {
	res := 0
	for i := 0; i < 64; i++ {
		res += int(x & 1)
		x >>= 1
	}
	return res
}

func main() {
	fmt.Println(popCountV2(25332))
}
