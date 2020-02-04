package main

import (
	"fmt"
	"math/rand"
)

func merge(a []int, b []int) {
	var c []int
	i, j := 0, 0
	lengthA, lengthB := len(a), len(b)

	for i < lengthA && j < lengthB {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	if i == lengthA {
		for j < lengthB {
			c = append(c, b[j])
			j++
		}
	} else {
		for i < lengthA {
			c = append(c, a[i])
			i++
		}
	}
	copy(a, c[:lengthA]) //copy: 拷贝同类型对slice，返回拷贝对元素数，是dst和src长度对较小值
	copy(b, c[lengthA:])
}

func mergeSort(a []int) {
	length := len(a)
	if length > 1 {
		part1 := a[:length/2]
		part2 := a[length/2:]
		mergeSort(part1)
		mergeSort(part2)
		merge(part1, part2)
	}
}

func main() {
	var a [100]int
	for i := range a {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	mergeSort(a[:])
	fmt.Println(a)
}
