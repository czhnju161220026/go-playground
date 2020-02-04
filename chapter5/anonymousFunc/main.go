package main

import "fmt"

import "strings"

func main() {
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-90000"))
}
