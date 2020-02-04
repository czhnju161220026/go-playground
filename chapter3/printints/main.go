package main

import "bytes"

import "fmt"

func ints2string(a []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range a {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(ints2string([]int{1, 2, 3, 4}))
}
