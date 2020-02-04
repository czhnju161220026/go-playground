package main

import "math"

import "fmt"

type Point3D struct {
	x, y, z float64
}

//由于主调函数会复制每一个实参变量，为了避免复制，可以用指针, 这也适用于更新接收者
func (p *Point3D) ScaleBy(s float64) {
	p.x *= s
	p.y *= s
	p.z *= s
}

func (p *Point3D) Distance(q *Point3D) float64 {
	return math.Sqrt((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y) + (p.z-q.z)*(p.z-q.z))
}

func main() {
	p := Point3D{1, 2, 3}
	q := Point3D{4, 5, 6}
	p.ScaleBy(1.5)
	fmt.Println(p)
	fmt.Println(p.Distance(&q))

}
