package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	s := "Hello world"
	var c ByteCounter
	fmt.Fprintf(&c, "%s", s)
	fmt.Println(c)
}
