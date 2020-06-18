package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/mattdavis0351/hub-me/icons"
	"github.com/mattdavis0351/hub-me/items"
	"github.com/mattdavis0351/hub-me/repository"
	"github.com/mattdavis0351/hub-me/utils"
)

func main() {
	fmt.Println(os.Getenv("HUB_ME_TOKEN"))
	systray.Run(entry, exit)
}

func entry() {
	// systray.SetTitle("Hub Me")
	systray.SetTooltip("GitHub Integration")
	systray.SetIcon(icons.Data)

	// Layout potential options
	mNewRepo := systray.AddMenuItem("New Repo", "Open browser to create new repo")
	systray.AddSeparator()

	issues := systray.AddMenuItem("GitHub Issues", "This is a list of issues that could use a better tooltip")

	issuesAssigned := issues.AddSubMenuItem("Assigned Issues", "View a list of assigned issues")
	utils.AddGitHubItems(issuesAssigned, "issue", "assigned")

	issuesMentioned := issues.AddSubMenuItem("Mentioned Issues", "View issues where you were mentioned")
	utils.AddGitHubItems(issuesMentioned, "issue", "mentioned")

	issuesCreated := issues.AddSubMenuItem("Created Issues", "View issues that you created")
	utils.AddGitHubItems(issuesCreated, "issue", "created")

	pulls := systray.AddMenuItem("GitHub Pull Requests", "This is a list of PRs that could use a better tooltip")

	pullsAssigned := pulls.AddSubMenuItem("Assigned Pull Requests", "View a list of assigned pulls")
	utils.AddGitHubItems(pullsAssigned, "pulls", "assigned")

	pullsMentioned := pulls.AddSubMenuItem("Mention Pull Requests", "")
	utils.AddGitHubItems(pullsMentioned, "pulls", "mentioned")

	pullsCreated := pulls.AddSubMenuItem("Creatd Pull Requests", "")
	utils.AddGitHubItems(pullsCreated, "pulls", "created")

	// TODO... implement review requested
	// pullReviewRequested := pulls.AddSubMenuItem("Review Requested", "")
	// utils.AddGitHubItems(pullReviewRequested, "pulls", "subcribed")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Exit the application")

	// Add click listeners to desired options
	go items.AppQuit(mQuit)
	go repository.NewRepo(mNewRepo)

}

func exit() {
	fmt.Println("program exited")
}
