package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createAddFilePage() fyne.CanvasObject {
	pageLangMap := langMap["addFilePage"][currentLanguage]

	topBar := createTopBar(pageLangMap["pageTitle"], createHomePage)

	fileNameEntry := widget.NewEntry()
	openFileButton := createFileOpenButton(pageLangMap["openFileButton"])

	neededDateButton := createCalendarButton(langMap["addFilePage"][currentLanguage]["neededDateButton"])
	startDateButton := createCalendarButton(langMap["addFilePage"][currentLanguage]["startDateButton"])
	endDateButton := createCalendarButton(langMap["addFilePage"][currentLanguage]["endDateButton"])
	locationEntry := widget.NewEntry()
	descriptionEntry := widget.NewMultiLineEntry()

	form := widget.NewForm(
		&widget.FormItem{
			Text:     pageLangMap["fileNameLabel"],
			Widget:   fileNameEntry,
			HintText: pageLangMap["fileNameHintText"],
		},
		&widget.FormItem{
			Text:     pageLangMap["openFileLabel"],
			Widget:   openFileButton,
			HintText: pageLangMap["openFileHintText"],
		},
		&widget.FormItem{
			Text:     "",
			Widget:   widget.NewLabel(""),
			HintText: "",
		},
		&widget.FormItem{
			Text:     pageLangMap["neededDateLabel"],
			Widget:   neededDateButton,
			HintText: pageLangMap["neededDateHintText"],
		},
		&widget.FormItem{
			Text:     pageLangMap["startDateLabel"],
			Widget:   startDateButton,
			HintText: pageLangMap["startDateHintText"],
		},
		&widget.FormItem{
			Text:     pageLangMap["endDateLabel"],
			Widget:   endDateButton,
			HintText: pageLangMap["endDateHintText"],
		},
		&widget.FormItem{
			Text:     pageLangMap["locationLabel"],
			Widget:   locationEntry,
			HintText: pageLangMap["locationHintText"],
		},
		&widget.FormItem{
			Text:     pageLangMap["descriptionLabel"],
			Widget:   descriptionEntry,
			HintText: pageLangMap["descriptionHintText"],
		},
	)

	form.OnSubmit = func() {}
	form.SubmitText = pageLangMap["formSubmit"]

	formWrapper := container.NewCenter(form)

	return container.NewBorder(topBar, nil, nil, nil, formWrapper)
}
