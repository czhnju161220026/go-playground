package main

import "bytes"

import "fmt"

func addComma2(a string) string {
	var buf bytes.Buffer
	length := len(a)
	for i, s := range a {
		if i != 0 && (length-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(byte(s))
	}
	return buf.String()
}

func main() {
	fmt.Println(addComma2("12387129461926491273981"))
	fmt.Println(addComma2("123"))
	fmt.Println(addComma2("1"))
}
