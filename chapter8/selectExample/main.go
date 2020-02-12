package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
	select{
	case ch<-i:
		fmt.Println("The channel is empty, Send",i)
	case x:=<-ch:
		fmt.Println(x)
	}
	}

}