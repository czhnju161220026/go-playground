package lengthConv

type Meter float64
type Inch float64

func ItoM(l Inch) Meter {
	return Meter(l / 40.0)
}

func MtoI(l Meter) Inch {
	return Inch(l * 40.0)
}
