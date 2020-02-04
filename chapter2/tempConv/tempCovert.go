package tempConv

type Celsius float64
type Fahrenheit float64
type K float64

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

func KtoC(t K) Celsius {
	return Celsius(t + K(AbsoluteZero))
}

func CtoK(t Celsius) K {
	return K(t - AbsoluteZero)
}
