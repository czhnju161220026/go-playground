package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Bad Args")
	}

	s := os.Args[1]
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		s = s[0:dot]
	}
	fmt.Println(s)

}
