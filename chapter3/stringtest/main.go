package main

import "fmt"

func main() {
	fmt.Print("Hello world \n")
	fmt.Print(`Hello world \n`) //原生的字符串字面量不会进行转义
}
