package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createChangePasswordPage() fyne.CanvasObject {
	pageLangMap := langMap["changePasswordPage"][currentLanguage]

	topBar := createTopBar(pageLangMap["pageTitle"], createSettingsPage)

	oldPasswordEntry := widget.NewPasswordEntry()
	newPasswordEntry := widget.NewPasswordEntry()
	changePasswordForm := widget.NewForm(
		&widget.FormItem{
			Text:     pageLangMap["oldPasswordLabel"],
			Widget:   oldPasswordEntry,
			HintText: pageLangMap["oldPasswordHint"],
		},
		&widget.FormItem{
			Text:     pageLangMap["newPasswordLabel"],
			Widget:   newPasswordEntry,
			HintText: pageLangMap["newPasswordHint"],
		},
	)
	changePasswordForm.OnSubmit = func() { changePassword(oldPasswordEntry, newPasswordEntry) }

	pageContent := container.NewCenter(changePasswordForm)

	return container.NewBorder(topBar, nil, nil, nil, pageContent)
}
