package main

import "fmt"

func f(s []int) {
	for i := range s {
		s[i] = 0
	}
}

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	f(a[1:4]) //可以通过对slice对修改完成对底层数组对修改
	fmt.Println(a)
}
