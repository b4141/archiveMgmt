package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var App = app.New()
var window = App.NewWindow("ArchiveMgmt")

func main() {
	homePage := widget.NewLabel("Hello ArchiveMgmt")

	window.SetContent(homePage)
	windowSize := fyne.NewSize(650, 700)
	window.Resize(windowSize)
	window.ShowAndRun()
}
