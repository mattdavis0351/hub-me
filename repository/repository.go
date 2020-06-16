package repository

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/mattdavis0351/hub-me/utils"
)

// NewRepo needs a comment
func NewRepo(m *systray.MenuItem) {
	<-m.ClickedCh
	status := utils.OpenBrowser("https://github.com/new")
	if status != true {
		fmt.Println("something went wrong opening your browser.... use a Mac dummy")
	}

}
