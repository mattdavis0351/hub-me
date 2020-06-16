package utils

import (
	"fmt"

	"github.com/getlantern/systray"
	gh "github.com/mattdavis0351/hub-me/github"
)

// PopulateAssignedList needs a comment so you leave me alone
func PopulateAssignedList(m *systray.MenuItem) {

	assignedIssues, err := gh.GetAssignedIssues()
	if err != nil {
		fmt.Println((err))
	}

	for _, v := range assignedIssues {
		mitem := m.AddSubMenuItem(*v.Title, "")
		mitem.Show()
		go AddClickToSubMenu(v, mitem)

	}

}

// PopulateMentionedList needs a comment so you leave me alone
func PopulateMentionedList(m *systray.MenuItem) {

	assignedIssues, err := gh.GetMentionedIssues()
	if err != nil {
		fmt.Println((err))
	}

	for _, v := range assignedIssues {
		mitem := m.AddSubMenuItem(*v.Title, "")
		mitem.Show()
		go AddClickToSubMenu(v, mitem)

	}

}

// PopulateCreatedList needs a comment so you leave me alone
func PopulateCreatedList(m *systray.MenuItem) {

	assignedIssues, err := gh.GetCreatedIssues()
	if err != nil {
		fmt.Println((err))
	}

	for _, v := range assignedIssues {
		mitem := m.AddSubMenuItem(*v.Title, "")
		mitem.Show()
		go AddClickToSubMenu(v, mitem)

	}

}

// PopulateSubscribedList needs a comment so you leave me alone
func PopulateSubscribedList(m *systray.MenuItem) {

	assignedIssues, err := gh.GetSubscribedIssues()
	if err != nil {
		fmt.Println((err))
	}

	for _, v := range assignedIssues {
		mitem := m.AddSubMenuItem(*v.Title, "")
		mitem.Show()
		go AddClickToSubMenu(v, mitem)

	}

}
