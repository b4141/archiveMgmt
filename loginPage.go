package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func checkPassword(password string, feedbackLabel *canvas.Text) {
	pageLangMap := langMap["loginPage"][currentLanguage]

	if password == currentPassword {
		window.SetContent(createHomePage())
	} else {
		feedbackLabel.Text = pageLangMap["wrongPassword"]
	}
}

func createLoginPage() fyne.CanvasObject {
	passwordInput := widget.NewPasswordEntry()
	feedbackLabel := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	loginButton := widget.NewButtonWithIcon("", theme.LoginIcon(), func() { checkPassword(passwordInput.Text, feedbackLabel) })

	loginPageContainer := container.NewBorder(nil, nil, nil, loginButton, passwordInput)
	separator := layout.NewSpacer()
	emptyLabel := widget.NewLabel("\t\t\t")

	return container.NewVBox(
		separator,
		container.NewBorder(nil, nil, emptyLabel, emptyLabel, loginPageContainer),
		container.NewBorder(nil, nil, emptyLabel, emptyLabel, feedbackLabel),
		separator,
	)
}
