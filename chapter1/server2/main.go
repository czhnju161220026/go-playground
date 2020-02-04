package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func handleEcho(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count += 1
	_, _ = fmt.Fprintf(w, "URL = %s\n", r.URL.Path)
	mu.Unlock()
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", handleEcho)
	http.HandleFunc("/count", handleCount)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
