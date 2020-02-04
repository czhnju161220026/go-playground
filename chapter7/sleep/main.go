package main

import "flag"

import "time"

import "fmt"

var period = flag.Duration("period", time.Second*1, "Sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleep for %v ...", *period)
	time.Sleep(*period)
	fmt.Println()
}
