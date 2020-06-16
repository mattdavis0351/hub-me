package utils

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/google/go-github/github"
)

// AddClickToSubMenu needs a comment
func AddClickToSubMenu(i github.Issue, m *systray.MenuItem) {

	<-m.ClickedCh
	status := OpenBrowser(*i.HTMLURL)
	if status != true {
		fmt.Println("something went wrong opening your browser.... use a Mac dummy")
	}

}
