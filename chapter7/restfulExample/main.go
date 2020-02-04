package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type database struct {
	content map[string]float64
}

type pair struct {
	Name  string
	Price float64
}

func (db *database) get(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message":"get called",`)
	fmt.Fprint(w, `"content":{`)
	for item, price := range db.content {
		fmt.Fprintf(w, "{%q:\"%v\",},", item, price)
	}
	fmt.Fprint(w, "},}")
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(body))
	items := []pair{}
	err := json.Unmarshal(body, &items)
	fmt.Println(items)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	for _, item := range items {
		db.content[item.Name] = item.Price
	}
	fmt.Fprint(w, `{"message":"update called"}`)
}

func main() {
	content := map[string]float64{"shoe": 50, "sock": 5}
	db := database{
		content,
	}
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", db.get).Methods(http.MethodGet)
	api.HandleFunc("", db.update).Methods(http.MethodPost)
	api.HandleFunc("", db.update).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8080", r))
}
