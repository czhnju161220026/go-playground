package main

import (
	"fmt"
	"os"
)

func fib(n int) (int, error) {
	if n == 1 || n == 2 {
		return 1, nil
	}
	if n > 2 {
		res := 0
		f1, f2 := 1, 1
		for i := 2; i < n; i++ {
			res = f1 + f2
			f1 = f2
			f2 = res
		}
		return res, nil
	}

	return -1, fmt.Errorf("invalid input: %v\n", n)
}

func main() {
	for i := -5; i < 10; i++ {
		res, err := fib(i)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		} else {
			fmt.Printf("Fib(%v) = %v\n", i, res)
		}
	}
}
