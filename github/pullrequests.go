package github

import (
	"fmt"

	"github.com/google/go-github/github"
)

func GetAssignedPulls() ([]github.Issue, error) {

	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: "assigned", State: "open", ListOptions: *listOptions}

	issue, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var pulls []github.Issue

	for i := 0; i < len(issue); i++ {
		if issue[i].IsPullRequest() {
			pulls = append(pulls, *issue[i])
		}
	}

	return pulls, nil

}

func GetMentionedPulls() ([]github.Issue, error) {

	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: "mentioned", State: "open", ListOptions: *listOptions}

	issue, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var pulls []github.Issue

	for i := 0; i < len(issue); i++ {
		if issue[i].IsPullRequest() {
			pulls = append(pulls, *issue[i])
		}
	}

	return pulls, nil

}

func GetCreatedPulls() ([]github.Issue, error) {

	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: "created", State: "open", ListOptions: *listOptions}

	issue, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var pulls []github.Issue

	for i := 0; i < len(issue); i++ {
		if issue[i].IsPullRequest() {
			pulls = append(pulls, *issue[i])
		}
	}

	return pulls, nil

}

func GetAllPulls() ([]github.Issue, error) {

	listOptions := &github.ListOptions{PerPage: 100}
	opt := &github.IssueListOptions{Filter: "all", State: "open", ListOptions: *listOptions}

	issue, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		fmt.Println(err)
	}

	var pulls []github.Issue

	for i := 0; i < len(issue); i++ {
		if issue[i].IsPullRequest() {
			pulls = append(pulls, *issue[i])
		}
	}

	return pulls, nil

}
