package repository

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/getlantern/systray"
)

// NewRepo needs a comment
func NewRepo(m *systray.MenuItem) {
	<-m.ClickedCh
	status := openBrowser("https://github.com/new")
	if status != true {
		fmt.Println("something went wrong opening your browser.... use a Mac dummy")
	}

}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
