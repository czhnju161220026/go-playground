package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

func reverseString(s string) string {
	res := []rune(s)
	for from, to := 0, len(res)-1; from < to; from, to = from+1, to-1 {
		res[from], res[to] = res[to], res[from]
	}
	return string(res)
}

func addComma3(val float64) string {
	var buf bytes.Buffer
	if val < 0.0 {
		buf.WriteRune('-')
		val *= -1
	}
	s := strconv.FormatFloat(val, 'f', 10, 64)
	tokens := strings.Split(s, ".")
	s = addComma2(tokens[0]) + "." + reverseString(addComma2(reverseString(tokens[1])))
	return s
}

func main() {
	fmt.Println(addComma3(12432523.34124124))
	fmt.Println(addComma3(1.213124))
	fmt.Println(addComma3(0.213))
	fmt.Println(addComma3(124))
	fmt.Println(math.Floor(3.4))
}
