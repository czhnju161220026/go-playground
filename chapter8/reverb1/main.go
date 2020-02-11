package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, content string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(content))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", content)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(content))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1000*time.Millisecond)
	}
}

func listen() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func main() {
	listen()
}
