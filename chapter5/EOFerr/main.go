package main

import "bufio"

import "os"

import "io"

import "fmt"

func main() {
	in := bufio.NewReader(os.Stdin)
	res := []rune{}
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			fmt.Println("Meet EOF.")
			break
		}
		if err != nil {
			fmt.Println("failed")
			os.Exit(1)
		}
		res = append(res, r)
	}
	fmt.Println(string(res))
}
