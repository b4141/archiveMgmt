package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createChangePasswordPage() fyne.CanvasObject {
	currentPageLanguage := map[string]map[string]string{
		"en": {
			"pageTitle":        "Change Password",
			"oldPasswordLabel": "Old Password:",
			"newPasswordLabel": "New Password:",
			"oldPasswordHint":  "Enter your Old Password here.\t",
			"newPasswordHint":  "Enter your New Password here.\t",
		},
		"fr": {
			"pageTitle":        "changer le mot de passe",
			"oldPasswordLabel": "Ancien mot de passe:",
			"newPasswordLabel": "Nouveau mot de passe:",
			"oldPasswordHint":  "Entrez votre ancien mot de passe ici.\t",
			"newPasswordHint":  "Entrez votre nouveau mot de passe ici.\t",
		},
		"ar": {
			"pageTitle":        "تغير كلمة السر",
			"oldPasswordLabel": "كلمة السر القديمة :",
			"newPasswordLabel": "كلمة السر الجديدة :",
			"oldPasswordHint":  "ادخل كلمة السر الحالية هنا        ",
			"newPasswordHint":  "ادخل كلمة السر الجديدة هنا        ",
		},
	}

	pageTitle := widget.NewLabel(currentPageLanguage[currentLanguage]["pageTitle"])
	pageTitle.TextStyle = widget.RichTextStyleHeading.TextStyle
	returnButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		window.SetContent(createSettingsPage())
	})
	topBarContainer := container.NewBorder(nil, nil, pageTitle, returnButton, nil)

	oldPasswordEntry := widget.NewPasswordEntry()
	newPasswordEntry := widget.NewPasswordEntry()
	changePasswordForm := widget.NewForm(
		&widget.FormItem{
			Text:     currentPageLanguage[currentLanguage]["oldPasswordLabel"],
			Widget:   oldPasswordEntry,
			HintText: currentPageLanguage[currentLanguage]["oldPasswordHint"],
		},
		&widget.FormItem{
			Text:     currentPageLanguage[currentLanguage]["newPasswordLabel"],
			Widget:   newPasswordEntry,
			HintText: currentPageLanguage[currentLanguage]["newPasswordHint"],
		},
	)
	changePasswordForm.OnSubmit = func() { changePassword(oldPasswordEntry, newPasswordEntry) }

	pageContent := container.NewCenter(changePasswordForm)

	return container.NewBorder(topBarContainer, nil, nil, nil, pageContent)
}
