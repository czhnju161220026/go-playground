package main

import (
	"fmt"
	"os"
)

func addComma(s string) string {
	if len(s) < 3 {
		return s
	}
	return addComma(s[:len(s)-3]) + "," + s[len(s)-3:]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Bad args")
		return
	}
	fmt.Println(addComma(os.Args[1]))
}
