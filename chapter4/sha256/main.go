package main

import "crypto/sha256"

import "fmt"

func main() {
	c1 := sha256.Sum256([]byte{'x'})
	c2 := sha256.Sum256([]byte{'y'})
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
