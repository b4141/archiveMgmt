package main

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func setTheme(themeName string) {
	if themeName == "dark" {
		App.Settings().SetTheme(theme.DarkTheme())
		currentTheme = "dark"
	} else {
		App.Settings().SetTheme(theme.LightTheme())
		currentTheme = "light"
	}
}

func setLanguage(languageName string) {
	if languageName == "English" {
		currentLanguage = "en"
	} else if languageName == "Francais" {
		currentLanguage = "fr"
		// } else if languageName == "العربية" {
		// 	currentLanguage = "ar"
	} else {
		currentLanguage = "en"
	}
}

func changePassword(oldPassword, newPassword *widget.Entry) {
	if oldPassword.Text == currentPassword {
		println("your new password is", newPassword.Text)
	} else {
		println("old password wrong!!")
	}
}

func changeSaveFolder(path string) {
	currentSaveFolder = path
	println("changed the path name to" + path)
}
