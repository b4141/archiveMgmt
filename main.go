package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var App = app.New()
var window = App.NewWindow("ArchiveMgmt")

func main() {
	// homePage := createLoginPage()
	homePage := createChangePasswordPage()

	window.SetContent(homePage)
	windowSize := fyne.NewSize(650, 700)
	window.Resize(windowSize)
	window.ShowAndRun()
}
