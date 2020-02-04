package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
var max = flag.Int("m", -1, "max words")

func main() {
	flag.Parse()
	if *max == -1 {
		fmt.Println(strings.Join(flag.Args(), *sep))
	} else {
		res := ""
		for i := 0; i < *max; i++ {
			res += flag.Arg(i) + *sep
		}
		fmt.Println(res)
	}

	if !*n {
		fmt.Println()
	}
}
