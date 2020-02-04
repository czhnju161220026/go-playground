package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	file, err := os.Open("chapter1/test.txt",)
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(file)
	counts := make(map[string]int)
	for input.Scan(){
		counts[input.Text()] += 1
	}
	for line, count := range counts {
		fmt.Printf("%s: %d\n", line, count)
	}

}
