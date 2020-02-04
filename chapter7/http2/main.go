package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollers float32

func (d dollers) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollers

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %s\n", item)
			return
		}
		fmt.Fprintf(w, "price: %s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", r.URL)
	}
}

func main() {
	db := database{"shoe": 50, "sock": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
