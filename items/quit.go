package items

import "github.com/getlantern/systray"

// AppQuit needs a comment
func AppQuit(m *systray.MenuItem) {
	<-m.ClickedCh
	systray.Quit()
}
