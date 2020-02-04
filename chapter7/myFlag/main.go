package main

import "strconv"

import "flag"

import "fmt"

type HexValue struct {
	val int
}

func (h *HexValue) Set(s string) error {
	var err error
	h.val, err = strconv.Atoi(s)
	return err
}

func (h *HexValue) String() string {
	res := []byte{}
	temp := h.val
	for temp != 0 {
		next := temp % 16
		temp /= 16
		if next <= 9 {
			res = append(res, byte(next)+'0')
		} else {
			res = append(res, byte(next-10)+'A')
		}
	}
	res = append(res, 'x', '0')
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func NewHexFlag(name string, value int, usage string) *HexValue {
	f := HexValue{val: value}
	flag.CommandLine.Var(&f, name, usage)
	return &f
}

var hexValue = NewHexFlag("value", 20, "Parse a Dec to Hex")

func main() {
	flag.Parse()
	fmt.Println(hexValue)
}
