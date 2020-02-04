package main

import "fmt"

//call by value
func f(a [6]int) {
	for i := range a {
		a[i] = 0
	}
}

//call by reference
func g(a *[6]int) {
	for i := range a {
		a[i] = 0
	}
}

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a=%v\n", a)
	f(a)
	fmt.Printf("a=%v\n", a)
	g(&a)
	fmt.Printf("a=%v\n", a)
}
