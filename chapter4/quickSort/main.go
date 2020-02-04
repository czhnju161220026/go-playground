package main

import "math/rand"

import "fmt"

func partition(a []int) int {
	pivot := a[0]
	splitPoint := 0
	for i := range a {
		if a[i] < pivot {
			splitPoint++
			if i != splitPoint {
				a[i], a[splitPoint] = a[splitPoint], a[i]
			}
		}
	}
	a[0], a[splitPoint] = a[splitPoint], a[0]
	return splitPoint
}

func qsort(a []int) {
	if len(a) > 1 {
		splitPoint := partition(a)
		//fmt.Println("Split point is: ", splitPoint)
		//fmt.Printf("Next part1: %v, part2: %v\n", a[:splitPoint], a[splitPoint:])
		qsort(a[:splitPoint])
		if splitPoint+1 < len(a)-1 {
			qsort(a[splitPoint+1:])
		}
	}
}

func main() {
	var a [100]int
	for i := range a {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	qsort(a[:])
	fmt.Println(a)
}
