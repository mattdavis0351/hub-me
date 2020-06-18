package utils

import (
	"fmt"

	"github.com/getlantern/systray"
)

func AddGitHubItems(m *systray.MenuItem, t, f string) {
	items, err := GetGitHubItems(t, f)
	if err != nil {
		fmt.Println((err))
	}

	for _, v := range items {
		mitem := m.AddSubMenuItem(*v.Title, "")
		go AddClickToSubMenu(v, mitem)

	}
}
