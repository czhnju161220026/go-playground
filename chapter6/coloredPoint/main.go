package main

import "image/color"

import "math"

import "fmt"

type Point2D struct {
	X, Y float64
}

func (this *Point2D) Distance(p Point2D) float64 {
	return math.Hypot(this.X-p.X, this.Y-p.Y)
}

type ColoredPoint2D struct {
	Point2D
	Color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint2D{Point2D{0, 0}, red}
	q := ColoredPoint2D{Point2D{9, 10}, blue}
	fmt.Println(p.Distance(q.Point2D))
}
