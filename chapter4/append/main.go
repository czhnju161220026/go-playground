package main

import "fmt"

func appendInt(a []int, b int) []int {
	var z []int
	zlen := len(a) + 1
	if zlen <= cap(a) {
		z = a[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(a) {
			zcap = 2 * len(a)
		}
		z = make([]int, zlen, zcap)
		copy(z, a)
	}
	z[len(a)] = b
	return z
}

func main() {
	a := []int{1}
	fmt.Printf("%d\t%d\t%v\n", cap(a), len(a), a)
	for i := 2; i < 10; i++ {
		a = appendInt(a, i)
		fmt.Printf("%d\t%d\t%v\n", cap(a), len(a), a)
	}
}
