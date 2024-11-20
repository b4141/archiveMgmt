package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func checkPassword(password string) {
	if password == currentPassword {
		println("correct password")
	} else {
		println("wrong password")
	}

}

func createLoginPage() fyne.CanvasObject {
	passwordInput := widget.NewPasswordEntry()
	loginButton := widget.NewButtonWithIcon("", theme.LoginIcon(), func() { checkPassword(passwordInput.Text) })

	loginPageContainer := container.NewBorder(nil, nil, nil, loginButton, passwordInput)
	separator := layout.NewSpacer()
	emptyLabel := widget.NewLabel("\t\t\t")

	return container.NewVBox(
		separator,
		container.NewBorder(nil, nil, emptyLabel, emptyLabel, loginPageContainer),
		separator,
	)
}
