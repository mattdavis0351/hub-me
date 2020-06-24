package items

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/mattdavis0351/hub-me/utils"
)

// AppLogin needs a comment
func AppLogin(m *systray.MenuItem) {
	<-m.ClickedCh

	status := utils.OpenBrowser("http://localhost:3000/login")
	if status != true {
		fmt.Println("something went wrong opening your browser.... use a Mac dummy")
	}

}
