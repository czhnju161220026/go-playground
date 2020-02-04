package main

import (
	"fmt"
	"os"
)

func main(){
	sep, s := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
