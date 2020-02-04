package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()] += 1
	}
}

func main()  {
	counts := make(map[string]int)
	fileLists := [1]string {"chapter1/test.txt"}
	for _, f := range fileLists{
		file, err := os.Open(f)
		if err != nil{
			log.Fatal(err)
		}
		countLines(file, counts)
	}

	for line,count := range counts{
		fmt.Printf("%s:%d\n", line, count)
	}
}
