package main

import "math"

import "fmt"

type Point2D struct {
	x, y float64
}

//普通的函数
func Distance(p, q Point2D) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

//Point2D类型的方法
func (this Point2D) Distance(p Point2D) float64 {
	return math.Hypot(this.x-p.x, this.y-p.y)
}

type Path []Point2D

//可以将方法绑定到任何类型上
func (path Path) Length() float64 {
	res := 0.0
	start := path[0]
	for _, p := range path[1:] {
		res += p.Distance(start)
		start = p
	}
	return res
}

func main() {
	p := Point2D{x: 0, y: 0}
	q := Point2D{x: 3, y: 4}
	path := Path{p, q}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Println(path.Length())
}
