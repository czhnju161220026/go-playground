package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(len(a))

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	b := [3]int{1, 2, 3}
	fmt.Println(b)
	c := [...]int{2, 3, 4}
	fmt.Println(c)
	d := [...]string{1: "first", 2: "second", 4: "forth"}
	fmt.Println(d)
	e := [...]int{10: -1}
	fmt.Println(e)
	f := [3]int{2, 3, 4}
	fmt.Println(b == c, c == f, f == b)
	g := []int{1, 2, 3}
	fmt.Printf("%T\n", g) //g是一个slice
	h := [...]int{1, 2, 3}
	fmt.Printf("%T\n", h) //h是一个array
}
