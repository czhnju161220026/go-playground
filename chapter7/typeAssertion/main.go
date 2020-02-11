package main

import (
	"fmt"
	"io"
	"sort"
)

type A interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
	F() int
}

type B struct{}

func (b *B) Len() int           { return 0 }
func (b *B) Swap(i, j int)      {}
func (b *B) Less(i, j int) bool { return false }
func (b *B) F() int             { return 0 }

func main() {
	var a A
	a = &B{}
	fmt.Println(a.(sort.Interface)) //成功
	fmt.Println(a.(io.ReadWriter))  //失败：缺少Read方法
}
