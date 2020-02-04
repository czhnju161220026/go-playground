package main

import (
	"fmt"
)

func f2c(t float64) float64 {
	return (t-35) * 5 / 9
}

func main() {
	fmt.Printf("%.4f F = %.4f C", 46.0, f2c(46))
}
