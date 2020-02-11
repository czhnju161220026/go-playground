package main

import "fmt"

func f(x interface{}) {
	switch x.(type) {
	case nil:
		fmt.Println("Nil")
	case int:
		fmt.Printf("Int: %d\n", x)
	case float64:
		fmt.Printf("Float64: %f\n", x)
	case string:
		fmt.Printf("String: %s\n", x)
	default:
		fmt.Println(x)
	}
}

func main() {
	f(nil)
	f(1234)
	f(12.5)
	f("Hello world")
}
