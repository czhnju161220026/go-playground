package main

import (
	"flag"
	"fmt"
	lc "goPlayground/chapter2/lengthConv"
	tc "goPlayground/chapter2/tempConv"
)

var flagT = flag.Bool("t", false, "Use temp convert(Default C to F)")
var flagL = flag.Bool("l", false, "Use length convert(Default meter to inch)")
var flagI = flag.Bool("i", false, "Inch to meter")
var flagF = flag.Bool("f", false, "F to C")
var num = flag.Float64("n", 0, "data")

func main() {
	flag.Parse()
	if !*flagT && !*flagL {
		fmt.Println("wrong flag")
		return
	} else if *flagT {
		if *flagF {
			fmt.Println(tc.FtoC(tc.Fahrenheit(*num)))
			return
		} else {
			fmt.Println(tc.CtoF(tc.Celsius(*num)))
			return
		}
	} else {
		if *flagI {
			fmt.Println(lc.ItoM(lc.Inch(*num)))
			return
		} else {
			fmt.Println(lc.MtoI(lc.Meter(*num)))
			return
		}
	}
}
