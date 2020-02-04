package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main(){
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Fetch %s error", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("%s", b)


	}
}
