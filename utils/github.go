package utils

import (
	"fmt"

	"github.com/google/go-github/github"
)

func GetGitHubItems(t, f string) ([]github.Issue, error) {
	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: f, State: "open", ListOptions: *listOptions}

	issue, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var arr []github.Issue

	if t == "pulls" {
		for i := 0; i < len(issue); i++ {
			if issue[i].IsPullRequest() {
				arr = append(arr, *issue[i])
			}
		}
	} else if t == "issue" {
		for i := 0; i < len(issue); i++ {
			if !issue[i].IsPullRequest() {
				arr = append(arr, *issue[i])
			}
		}
	}

	return arr, nil
}
