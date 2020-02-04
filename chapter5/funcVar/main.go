package main

import "fmt"

func add(op1, op2 int) int {
	return op1 + op2
}

func sub(op1, op2 int) int {
	return op1 - op2
}

func mul(op1, op2 int) int {
	return op1 * op2
}

func div(op1, op2 int) int {
	return op1 / op2
}

func calc(op1, op2 int, f func(a, b int) int) int {
	if f != nil {
		return f(op1, op2)
	} else {
		return 0
	}
}

func main() {
	fmt.Println(calc(1, 1, add))
	fmt.Println(calc(3, 2, sub))
	fmt.Println(calc(4, 5, mul))
	fmt.Println(calc(9, 3, div))
}
