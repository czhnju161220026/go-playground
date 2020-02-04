package githubissue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// IssuesURL ...
const IssuesURL = "https://api.github.com/search/issues"

// User ...
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// Issue ...
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// IssuesSearchResult ...
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// SearchIssues ...
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get error: %v\n", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "http.Status not OK: %v\n", err)
		return nil, err
	}

	var result IssuesSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "json decode error: %v\n", err)
		return nil, err
	}
	resp.Body.Close()
	return &result, err
}
