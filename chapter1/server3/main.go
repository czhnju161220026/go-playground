package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_,_ = fmt.Fprintf(w, "URL: %s\n", r.URL.Path)
	for k,v := range r.Header{
		_,_ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error parse form: %q\n", err)
	}
	for k,v := range r.Form{
		fmt.Fprintf(w, "Form[%q] = %q\n", k ,v)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

