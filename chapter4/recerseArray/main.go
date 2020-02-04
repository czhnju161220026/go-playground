package main

import "fmt"

func reverse(a []int) {
	l := len(a)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(a[:])
	fmt.Println(a)
}
