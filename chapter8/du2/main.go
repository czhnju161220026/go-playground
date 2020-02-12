package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSizes chan int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1fGB\n", nfiles, float64(nbytes)/1e9)
}

var output = flag.Bool("v", false, "whether to print the log")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)

	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *output {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
	end := false
	for {
		select {
		case n, ok := <-fileSizes:
			if !ok {
				end = true
				break
			} else {
				nfiles++
				nbytes += n
			}
		case <-tick:
			{
				printDiskUsage(nfiles, nbytes)
			}
		}
		if end {
			break
		}
	}
	printDiskUsage(nfiles, nbytes)
}
