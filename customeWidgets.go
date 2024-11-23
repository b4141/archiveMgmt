package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

func createCalendarButton(buttonText string) fyne.CanvasObject {
	button := widget.NewButton(buttonText, nil)
	button.OnTapped = func() {
		d := &dialog.CustomDialog{}
		calendar := xwidget.NewCalendar(time.Now(), func(t time.Time) {
			button.SetText(t.Format("02-01-2006"))
			d.Hide()
		})
		d = dialog.NewCustom(
			langMap["calendarDialog"][currentLanguage]["title"],
			langMap["calendarDialog"][currentLanguage]["cancel"],
			calendar,
			window,
		)
		d.Show()
	}

	return button
}

func createFileOpenButton(buttonText string) fyne.CanvasObject {
	button := widget.NewButton(buttonText, nil)
	button.OnTapped = func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
			}
			if reader == nil {
				println("canceled")
			}
			button.SetText(reader.URI().Name())
		}, window)
		fd.Show()
	}
	return button
}

func createTopBar(title string, content func() fyne.CanvasObject) fyne.CanvasObject {
	pageTitle := widget.NewLabel(title)
	pageTitle.TextStyle = widget.RichTextStyleHeading.TextStyle
	returnButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		if content != nil {
			window.SetContent(content())
		}
	})
	topBarContainer := container.NewBorder(nil, nil, pageTitle, returnButton, nil)

	return topBarContainer
}
