package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createChangePasswordPage() fyne.CanvasObject {
	pageTitle := widget.NewLabel(langMap["changePasswordPage"][currentLanguage]["pageTitle"])
	pageTitle.TextStyle = widget.RichTextStyleHeading.TextStyle
	returnButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		window.SetContent(createSettingsPage())
	})
	topBarContainer := container.NewBorder(nil, nil, pageTitle, returnButton, nil)

	oldPasswordEntry := widget.NewPasswordEntry()
	newPasswordEntry := widget.NewPasswordEntry()
	changePasswordForm := widget.NewForm(
		&widget.FormItem{
			Text:     langMap["changePasswordPage"][currentLanguage]["oldPasswordLabel"],
			Widget:   oldPasswordEntry,
			HintText: langMap["changePasswordPage"][currentLanguage]["oldPasswordHint"],
		},
		&widget.FormItem{
			Text:     langMap["changePasswordPage"][currentLanguage]["newPasswordLabel"],
			Widget:   newPasswordEntry,
			HintText: langMap["changePasswordPage"][currentLanguage]["newPasswordHint"],
		},
	)
	changePasswordForm.OnSubmit = func() { changePassword(oldPasswordEntry, newPasswordEntry) }

	pageContent := container.NewCenter(changePasswordForm)

	return container.NewBorder(topBarContainer, nil, nil, nil, pageContent)
}
