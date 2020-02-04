package main

import "net/url"

import "fmt"

func main() {
	m := url.Values{"lang": {"en", "zh"}}
	m.Add("item", "1")
	m.Add("item", "2")
	m.Add("lang", "jp")
	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["lang"])
	
}
