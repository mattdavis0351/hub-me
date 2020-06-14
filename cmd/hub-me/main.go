package main

import (
	"github.com/getlantern/systray"
	"github.com/mattdavis0351/hub-me/items"
	"github.com/mattdavis0351/hub-me/repository"
)

func main() {
	systray.Run(entry, exit)
}

func entry() {
	systray.SetTitle("Hub Me")
	systray.SetTooltip("GitHub Integration")
	mNewRepo := systray.AddMenuItem("New Repo", "Open browser to create new repo")
	mQuit := systray.AddMenuItem("Quit", "Exit the application")
	mNewRepo.Show()
	mQuit.Show()
	go items.AppQuit(mQuit)
	go repository.NewRepo(mNewRepo)

}

func exit() {
	panic("something broke")
}
