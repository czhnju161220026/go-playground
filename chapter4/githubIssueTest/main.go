package main

import (
	"fmt"
	gitissue "goPlayground/chapter4/githubissue"
	"os"
)

func main() {
	result, err := gitissue.SearchIssues(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
