package main

import "fmt"

// Point has x and y
type Point struct {
	X int
	Y int
}

// Circle is a circle
type Circle struct {
	Point
	Radius int
}

// Wheel is wheel
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w1 := Wheel{
		Circle: Circle{
			Point: Point{
				X: 10,
				Y: 20,
			},
			Radius: 15,
		},
		Spokes: 8,
	}

	w2 := Wheel{Circle{Point{10, 20}, 15}, 8}

	fmt.Printf("%#v\n", w1)
	fmt.Printf("%#v\n", w2)
}
