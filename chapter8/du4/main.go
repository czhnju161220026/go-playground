package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func dirents(dir string) []os.FileInfo {
	select {
	case <-done:
		return nil
	case tokens <- 1:
	}
	defer func() { <-tokens }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1fGB\n", nfiles, float64(nbytes)/1e9)
}

var output = flag.Bool("v", false, "whether to print the log")

var tokens = make(chan int, 20)
var done = make(chan int)

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)

	n := new(sync.WaitGroup)

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var tick <-chan time.Time
	if *output {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
	end := false
	for {
		select {
		case <-done:
			for range fileSizes {
				// do nothing
			}
			return
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
