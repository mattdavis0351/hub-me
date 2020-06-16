package github

import (
	"fmt"

	"github.com/google/go-github/github"
)

func GetAssignedIssues() ([]github.Issue, error) {

	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: "assigned", State: "open", ListOptions: *listOptions}

	issues, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var issue []github.Issue

	for i := 0; i < len(issues); i++ {
		if !issues[i].IsPullRequest() {
			issue = append(issue, *issues[i])
		}
	}

	return issue, nil
}

func GetMentionedIssues() ([]github.Issue, error) {

	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: "mentioned", State: "open", ListOptions: *listOptions}

	issues, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var issue []github.Issue

	for i := 0; i < len(issues); i++ {
		if !issues[i].IsPullRequest() {
			issue = append(issue, *issues[i])
		}
	}

	return issue, nil
}

func GetCreatedIssues() ([]github.Issue, error) {

	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: "created", State: "open", ListOptions: *listOptions}

	issues, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var issue []github.Issue

	for i := 0; i < len(issues); i++ {
		if !issues[i].IsPullRequest() {
			issue = append(issue, *issues[i])
		}
	}

	return issue, nil
}

func GetSubscribedIssues() ([]github.Issue, error) {

	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: "subscribed", State: "open", ListOptions: *listOptions}

	issues, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var issue []github.Issue

	for i := 0; i < len(issues); i++ {
		if !issues[i].IsPullRequest() {
			issue = append(issue, *issues[i])
		}
	}

	return issue, nil
}

func GetAllIssues() ([]github.Issue, error) {

	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: "subscribed", State: "open", ListOptions: *listOptions}

	issues, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var issue []github.Issue

	for i := 0; i < len(issues); i++ {
		if !issues[i].IsPullRequest() {
			issue = append(issue, *issues[i])
		}
	}

	return issue, nil
}
