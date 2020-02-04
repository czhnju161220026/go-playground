package main

import "os"

import "fmt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Bad args")
		return
	} 
	s := os.Args[1]
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '/' {
				s = s[i+1:]
				break
			}
		}

		for i := 0; i < len(s); i++ {
			if s[i] == '.' {
				s = s[0:i]
				break
			}
		}

		fmt.Println(s)
}
