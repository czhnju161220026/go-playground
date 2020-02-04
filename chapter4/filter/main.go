package main

import "fmt"

func filter(strs []string) []string {
	res := []string{}
	for _, str := range strs {
		if len(res) == 0 || res[len(res)-1] != str {
			res = append(res, str)
		}
	}
	return res
}

func main() {
	strs := []string{"hello", "hello", "java", "java", "go", "java", "c++", "c++", "python"}
	fmt.Println(filter(strs))
}
