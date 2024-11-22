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
	currentPageLanguage := map[string]map[string]string{
		"en": {
			"pageTitle":                  "Settings",
			"themeDark":                  "Dark Theme",
			"themeLight":                 "Light Theme",
			"themeLabel":                 "Theme:",
			"languageLabel":              "Language:",
			"changePasswordButton":       "Change Password",
			"changeSaveFolderButton":     "change save folder",
			"saveFolderPathLabelDeafult": "Deafult",
			"englishLanguage":            "English",
			"frenchLanguage":             "Francais",
			"arabicLanguage":             "العربية",
		},
		"fr": {
			"pageTitle":                  "FR_Settings",
			"themeDark":                  "FR_Dark Theme",
			"themeLight":                 "FR_Light Theme",
			"themeLabel":                 "FR_Theme:",
			"languageLabel":              "FR_Language:",
			"changePasswordButton":       "FR_Change Password",
			"changeSaveFolderButton":     "FR_change save folder",
			"saveFolderPathLabelDeafult": "FR_Deafult",
			"englishLanguage":            "English",
			"frenchLanguage":             "Francais",
			"arabicLanguage":             "العربية",
		},
		"ar": {
			"pageTitle":                  "الإعدادات",
			"themeDark":                  "مظلم",
			"themeLight":                 "مضيء",
			"themeLabel":                 "الوضع:",
			"languageLabel":              "اللغة:",
			"changePasswordButton":       "تغير كلمة السر",
			"changeSaveFolderButton":     "غير مجلد حفظ الملفات",
			"saveFolderPathLabelDeafult": "AR_Deafult",
			"englishLanguage":            "English",
			"frenchLanguage":             "Francais",
			"arabicLanguage":             "العربية",
		},
	}

	pageTitle := widget.NewLabel(currentPageLanguage[currentLanguage]["pageTitle"])
	pageTitle.TextStyle = widget.RichTextStyleHeading.TextStyle
	returnButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		println("exit")
	})
	topBarContainer := container.NewBorder(nil, nil, pageTitle, returnButton, nil)

	themeSelect := widget.NewSelect(
		[]string{
			currentPageLanguage[currentLanguage]["themeDark"],
			currentPageLanguage[currentLanguage]["themeLight"],
		},
		func(selectedTheme string) {
			if selectedTheme == currentPageLanguage[currentLanguage]["themeDark"] {
				setTheme("dark")
			} else {
				setTheme("light")
			}
		},
	)

	if currentTheme == "dark" {
		themeSelect.PlaceHolder = currentPageLanguage[currentLanguage]["themeDark"]
	} else if currentTheme == "light" {
		themeSelect.PlaceHolder = currentPageLanguage[currentLanguage]["themeLight"]
	}

	languageSelect := widget.NewSelect(
		[]string{
			currentPageLanguage[currentLanguage]["englishLanguage"],
			currentPageLanguage[currentLanguage]["frenchLanguage"],
			// currentPageLanguage[currentLanguage]["arabicLanguage"],
		},
		func(selectedLanguage string) {
			setLanguage(selectedLanguage)
			window.SetContent(createSettingsPage())
		},
	)

	if currentLanguage == "en" {
		languageSelect.PlaceHolder = currentPageLanguage[currentLanguage]["englishLanguage"]
	} else if currentLanguage == "fr" {
		languageSelect.PlaceHolder = currentPageLanguage[currentLanguage]["frenchLanguage"]
	} else if currentLanguage == "ar" {
		languageSelect.PlaceHolder = currentPageLanguage[currentLanguage]["arabicLanguage"]
	}

	changePassordButton := widget.NewButton(currentPageLanguage[currentLanguage]["changePasswordButton"], func() {
		window.SetContent(createChangePasswordPage())
	})

	changeSaveFolderButton := widget.NewButtonWithIcon(
		currentPageLanguage[currentLanguage]["changeSaveFolderButton"],
		theme.FolderIcon(),
		func() {
			openChangeSaveFolderDialog()
		},
	)

	saveFolderPathLabel := widget.NewLabel(currentPageLanguage[currentLanguage]["saveFolderPathLabelDeafult"])
	saveFolderPathLabel.TextStyle = widget.RichTextStyleBlockquote.TextStyle

	if currentSaveFolder != "./files" {
		splitedSFP := strings.Split(currentSaveFolder, "/")
		saveFolderPathLabel.SetText(".... / " + splitedSFP[len(splitedSFP)-2] + " / " + splitedSFP[len(splitedSFP)-1])
	}

	pageContent := container.NewCenter(container.NewVBox(
		widget.NewForm(
			&widget.FormItem{
				Text:     currentPageLanguage[currentLanguage]["themeLabel"],
				Widget:   themeSelect,
				HintText: "\t\t\t",
			},
			&widget.FormItem{
				Text:     currentPageLanguage[currentLanguage]["languageLabel"],
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
