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

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range *db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := (*db)[item]
	if !ok {
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func main() {
	db := database{"shoe": 50, "sock": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
