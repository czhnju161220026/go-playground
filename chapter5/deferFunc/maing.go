package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	log.Println("Enter: ", msg, "at", time.Now())
	return func() {
		log.Println("Leave: ", msg, "at", time.Now())
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(time.Second * 5)
}
func main() {
	bigSlowOperation()
}
