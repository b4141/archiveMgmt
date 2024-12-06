package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createHomePage() fyne.CanvasObject {
	pageLangMap := langMap["homePage"][currentLanguage]
	spacer := layout.NewSpacer()

	searchInput := widget.NewEntry()
	searchInput.SetPlaceHolder(pageLangMap["searchInputPlaceHolder"])

	addFileButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		window.SetContent(createAddFilePage())
	})
	settingsButton := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		window.SetContent(createSettingsPage())
	})
	logoutButton := widget.NewButtonWithIcon("", theme.LogoutIcon(), func() {
		window.SetContent(createLoginPage())
	})

	homeTopBarButtons := container.NewHBox(addFileButton, settingsButton, spacer, logoutButton)
	topBar := container.NewBorder(nil, container.NewHBox(spacer), nil, homeTopBarButtons, searchInput)

	pageContent := container.NewCenter(widget.NewLabel("page content goes here"))

	return container.NewBorder(topBar, nil, nil, nil, pageContent)
}
