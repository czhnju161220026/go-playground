package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Celsius = -273.5
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func CtoF(t Celsius) Fahrenheit {
	return Fahrenheit(t*5/9 + 32)
}

func FtoC(t Fahrenheit) Celsius {
	return Celsius((t - 32) * 9 / 5)
}

func main() {
	fmt.Printf("%.2f C = %.2f F", 30.0, CtoF(30.0))
}
