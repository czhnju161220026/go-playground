package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var contentPtr = flag.String("c", "null", "Give the content to be crypt")
var alg = flag.String("m", "sha256", "Choose the crypton algorithm")

func main() {
	flag.Parse()
	content := []byte(*contentPtr)
	if *alg == "sha256" {
		fmt.Printf("sha256(%v) = %x\n", *contentPtr, sha256.Sum256(content))
	} else if *alg == "sha384" {
		fmt.Printf("sha384(%v) = %x\n", *contentPtr, sha512.Sum384(content))
	} else if *alg == "sha512" {
		fmt.Printf("sha512(%v) = %x\n", *contentPtr, sha512.Sum512(content))
	}
}
