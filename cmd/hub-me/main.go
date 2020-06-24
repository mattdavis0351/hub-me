package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/getlantern/systray"
	"github.com/mattdavis0351/hub-me/auth"
	"github.com/mattdavis0351/hub-me/icons"
	"github.com/mattdavis0351/hub-me/items"
	"github.com/mattdavis0351/hub-me/repository"
	"github.com/mattdavis0351/hub-me/utils"
)

func main() {

	systray.Run(entry, exit)
}

func entry() {
	fmt.Println("systray is trash, this is the entrypoint")
	// systray.SetTitle("Hub Me")
	systray.SetTooltip("GitHub Integration")
	systray.SetIcon(icons.Data)

	for auth.UserLogin != true {
		fmt.Println("user login is not true...")
		mLogin := systray.AddMenuItem("Login", "")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Exit the application")
		go items.AppQuit(mQuit)
		go items.AppLogin(mLogin)

		http.HandleFunc("/login", auth.HandleGitHubLogin)
		http.HandleFunc("/login/cb", auth.HandleGitHubCallback)

		log.Fatal(http.ListenAndServe(":3000", nil))

		fmt.Println("while user login is not true....")
	}
	fmt.Println("user has logged in, you should see the new menu")
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

	// Add click listeners to desired options

	go repository.NewRepo(mNewRepo)

	mQuit := systray.AddMenuItem("Quit", "Exit the application")
	go items.AppQuit(mQuit)

}

func exit() {
	fmt.Println("program exited")
}
