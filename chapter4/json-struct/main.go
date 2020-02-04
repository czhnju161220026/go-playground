package main

import "encoding/json"

import "fmt"

import "os"

// Movie struct describe file
type Movie struct {
	Title  string
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

// MovieWithoutColor ...
type MovieWithoutColor struct {
	Title  string
	Year   int      `json:"released"`
	Actors []string `json:"actors"`
}

func main() {
	var movies = []Movie{
		{Title: "Movie1", Year: 1998, Color: false, Actors: []string{"Actor1", "Actor2"}},
		{Title: "Movie2", Year: 1988, Color: true, Actors: []string{"Actor1", "Actor2"}},
		{Title: "Movie3", Year: 2019, Color: true, Actors: []string{"Actor1", "Actor2"}},
	}

	jsonContent, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Json marsha error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", jsonContent)

	var data = `[
		{
		  "Title": "Movie1",
		  "released": 1998,
		  "actors": [
			"Actor1",
			"Actor2"
		  ]
		},
		{
		  "Title": "Movie2",
		  "released": 1988,
		  "color": true,
		  "actors": [
			"Actor1",
			"Actor2"
		  ]
		},
		{
		  "Title": "Movie3",
		  "released": 2019,
		  "color": true,
		  "actors": [
			"Actor1",
			"Actor2"
		  ]
		}
	  ]`

	var movies2 = []MovieWithoutColor{}
	err = json.Unmarshal([]byte(data), &movies2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Json unmarshal error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(movies2)
}
