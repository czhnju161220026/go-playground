package main

import (
	"fmt"
	tc "goPlayground/chapter2/tempConv" //需要是GoPath下的路径
	//可以给包绑定一个短名字，类似于python的 import xxx as x
)

func main() {
	fmt.Println(tc.AbsoluteZero)
	fmt.Println(tc.KtoC(0.0))
}
