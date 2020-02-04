package main

import "time"

import "text/tabwriter"

import "os"

import "fmt"

import "sort"

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byYear []Track

func (tracks byYear) Len() int {
	return len(tracks)
}

func (tracks byYear) Less(i, j int) bool {
	return tracks[i].Year < tracks[j].Year
}

func (tracks byYear) Swap(i, j int) {
	tracks[i], tracks[j] = tracks[j], tracks[i]
}

type byArtist []Track

func (tracks byArtist) Len() int {
	return len(tracks)
}

func (tracks byArtist) Less(i, j int) bool {
	return tracks[i].Artist < tracks[j].Artist
}

func (tracks byArtist) Swap(i, j int) {
	tracks[i], tracks[j] = tracks[j], tracks[i]
}

func printTracks(tracks []Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Duration")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "--------")
	for _, track := range tracks {
		fmt.Fprintf(tw, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}
	tw.Flush()
}

func main() {
	tracks := []Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m18s")},
		{"Go", "MOby", "Moby", 1992, length("1m20s")},
		{"Go Ahead", "Alica Keys", "As I am", 1999, length("4m36s")},
		{"Ready 2 Go", "Martin", "Smash", 2011, length("3m50s")},
	}
	printTracks(tracks)
	fmt.Println()
	sort.Sort(byYear(tracks))
	printTracks(tracks)
	fmt.Println()
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
}
