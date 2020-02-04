package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	// for i := range pc {
	// 	fmt.Println(pc[i])
	// }
}

func popCountV1(x uint64) int {
	return int(pc[byte(x)]) +
		int(pc[byte(x>>8)]) +
		int(pc[byte(x>>16)]) +
		int(pc[byte(x>>24)]) +
		int(pc[byte(x>>32)]) +
		int(pc[byte(x>>40)]) +
		int(pc[byte(x>>48)]) +
		int(pc[byte(x>>56)])
}

func main() {
	n := popCountV1(4235)
	fmt.Println(n)
}
