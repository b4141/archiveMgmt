package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func openChangeSaveFolderDialog() {
	fd := dialog.NewFolderOpen(
		func(lu fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if lu == nil {
				println("Cancelled")
				return
			}
			changeSaveFolder(lu.Path())
			window.SetContent(createSettingsPage())
		},
		window,
	)
	fd.Show()
}

func createSettingsPage() fyne.CanvasObject {
	pageTitle := widget.NewLabel(langMap["settingsPage"][currentLanguage]["pageTitle"])
	pageTitle.TextStyle = widget.RichTextStyleHeading.TextStyle
	returnButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		println("exit")
	})
	topBarContainer := container.NewBorder(nil, nil, pageTitle, returnButton, nil)

	themeSelect := widget.NewSelect(
		[]string{
			langMap["settingsPage"][currentLanguage]["themeDark"],
			langMap["settingsPage"][currentLanguage]["themeLight"],
		},
		func(selectedTheme string) {
			if selectedTheme == langMap["settingsPage"][currentLanguage]["themeDark"] {
				setTheme("dark")
			} else {
				setTheme("light")
			}
		},
	)

	if currentTheme == "dark" {
		themeSelect.PlaceHolder = langMap["settingsPage"][currentLanguage]["themeDark"]
	} else if currentTheme == "light" {
		themeSelect.PlaceHolder = langMap["settingsPage"][currentLanguage]["themeLight"]
	}

	languageSelect := widget.NewSelect(
		[]string{
			langMap["settingsPage"][currentLanguage]["englishLanguage"],
			langMap["settingsPage"][currentLanguage]["frenchLanguage"],
			// langMap["settingsPage"][currentLanguage]["arabicLanguage"],
		},
		func(selectedLanguage string) {
			setLanguage(selectedLanguage)
			window.SetContent(createSettingsPage())
		},
	)

	if currentLanguage == "en" {
		languageSelect.PlaceHolder = langMap["settingsPage"][currentLanguage]["englishLanguage"]
	} else if currentLanguage == "fr" {
		languageSelect.PlaceHolder = langMap["settingsPage"][currentLanguage]["frenchLanguage"]
	} else if currentLanguage == "ar" {
		languageSelect.PlaceHolder = langMap["settingsPage"][currentLanguage]["arabicLanguage"]
	}

	changePassordButton := widget.NewButton(langMap["settingsPage"][currentLanguage]["changePasswordButton"], func() {
		window.SetContent(createChangePasswordPage())
	})

	changeSaveFolderButton := widget.NewButtonWithIcon(
		langMap["settingsPage"][currentLanguage]["changeSaveFolderButton"],
		theme.FolderIcon(),
		func() {
			openChangeSaveFolderDialog()
		},
	)

	saveFolderPathLabel := widget.NewLabel(langMap["settingsPage"][currentLanguage]["saveFolderPathLabelDeafult"])
	saveFolderPathLabel.TextStyle = widget.RichTextStyleBlockquote.TextStyle

	if currentSaveFolder != "./files" {
		splitedSFP := strings.Split(currentSaveFolder, "/")
		saveFolderPathLabel.SetText(".... / " + splitedSFP[len(splitedSFP)-2] + " / " + splitedSFP[len(splitedSFP)-1])
	}

	pageContent := container.NewCenter(container.NewVBox(
		widget.NewForm(
			&widget.FormItem{
				Text:     langMap["settingsPage"][currentLanguage]["themeLabel"],
				Widget:   themeSelect,
				HintText: "\t\t\t",
			},
			&widget.FormItem{
				Text:     langMap["settingsPage"][currentLanguage]["languageLabel"],
				Widget:   languageSelect,
				HintText: "\t\t\t",
			},
		),
		widget.NewLabel(""),
		changePassordButton,
		widget.NewLabel(""),
		changeSaveFolderButton,
		saveFolderPathLabel,
	))

	return container.NewBorder(topBarContainer, nil, nil, nil, pageContent)
}
