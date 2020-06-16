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
	utils.PopulateAssignedList(issuesAssigned)
	issuesMentioned := issues.AddSubMenuItem("Mentioned Issues", "View issues where you were mentioned")
	utils.PopulateMentionedList(issuesMentioned)
	issuesCreated := issues.AddSubMenuItem("Created Issues", "View issues that you created")
	utils.PopulateMentionedList(issuesCreated)

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Exit the application")

	// Show all menu options
	mNewRepo.Show()
	issuesAssigned.Show()
	issuesMentioned.Show()
	mQuit.Show()

	// Add click listeners to desired options
	go items.AppQuit(mQuit)
	go repository.NewRepo(mNewRepo)

}

func exit() {
	fmt.Println("program exited")
}
