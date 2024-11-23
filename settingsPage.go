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
	pageLangMap := langMap["settingsPage"][currentLanguage]
	topBar := createTopBar(pageLangMap["pageTitle"], nil)

	themeSelect := widget.NewSelect(
		[]string{
			pageLangMap["themeDark"],
			pageLangMap["themeLight"],
		},
		func(selectedTheme string) {
			if selectedTheme == pageLangMap["themeDark"] {
				setTheme("dark")
			} else {
				setTheme("light")
			}
		},
	)

	if currentTheme == "dark" {
		themeSelect.PlaceHolder = pageLangMap["themeDark"]
	} else if currentTheme == "light" {
		themeSelect.PlaceHolder = pageLangMap["themeLight"]
	}

	languageSelect := widget.NewSelect(
		[]string{
			pageLangMap["englishLanguage"],
			pageLangMap["frenchLanguage"],
			// pageLangMap["arabicLanguage"],
		},
		func(selectedLanguage string) {
			setLanguage(selectedLanguage)
			window.SetContent(createSettingsPage())
		},
	)

	if currentLanguage == "en" {
		languageSelect.PlaceHolder = pageLangMap["englishLanguage"]
	} else if currentLanguage == "fr" {
		languageSelect.PlaceHolder = pageLangMap["frenchLanguage"]
	} else if currentLanguage == "ar" {
		languageSelect.PlaceHolder = pageLangMap["arabicLanguage"]
	}

	changePassordButton := widget.NewButton(pageLangMap["changePasswordButton"], func() {
		window.SetContent(createChangePasswordPage())
	})

	changeSaveFolderButton := widget.NewButtonWithIcon(
		pageLangMap["changeSaveFolderButton"],
		theme.FolderIcon(),
		func() {
			openChangeSaveFolderDialog()
		},
	)

	saveFolderPathLabel := widget.NewLabel(pageLangMap["saveFolderPathLabelDeafult"])
	saveFolderPathLabel.TextStyle = widget.RichTextStyleBlockquote.TextStyle

	if currentSaveFolder != "./files" {
		splitedSFP := strings.Split(currentSaveFolder, "/")
		saveFolderPathLabel.SetText(splitedSFP[0] + " / ... / " + splitedSFP[len(splitedSFP)-2] + " / " + splitedSFP[len(splitedSFP)-1])
	}

	pageContent := container.NewCenter(container.NewVBox(
		widget.NewForm(
			&widget.FormItem{
				Text:     pageLangMap["themeLabel"],
				Widget:   themeSelect,
				HintText: "\t\t\t",
			},
			&widget.FormItem{
				Text:     pageLangMap["languageLabel"],
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

	return container.NewBorder(topBar, nil, nil, nil, pageContent)
}
