package main

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createFileInfoPreview(filename string) fyne.CanvasObject {
	dumyText := ";alskdfj;alskfjasklalskfjasklalskfjaskl;\n;asldkjfasl;kfjasd\nlas;dkfjaskl;dfj\n;laksdjf;alksjfads\n;lkjasdf;klasjdfals;k\nasdl;kfjas;klfjds"
	spacer := widget.NewLabel("")
	formSpacer := &widget.FormItem{
		Text:   "",
		Widget: widget.NewLabel(""),
	}

	now := time.Now().Format(time.DateOnly)
	icon := widget.NewIcon(theme.FileIcon())
	fileName := widget.NewLabel(filename)
	nameBar := container.NewCenter(container.NewHBox(icon, fileName))

	descriptionScroll := container.NewScroll(widget.NewLabel(dumyText))
	descriptionScroll.SetMinSize(fyne.NewSize(120, 100))
	info := widget.NewForm(
		&widget.FormItem{
			Text:   "fileType:",
			Widget: widget.NewLabel("txt"),
		},
		&widget.FormItem{
			Text:   "added:",
			Widget: widget.NewLabel(now),
		},
		formSpacer,
		&widget.FormItem{
			Text:   "created Date:",
			Widget: widget.NewLabel(now),
		},
		&widget.FormItem{
			Text:   "expiration Date:",
			Widget: widget.NewLabel(now),
		},
		&widget.FormItem{
			Text:   "needed Date:",
			Widget: widget.NewLabel(now),
		},
		&widget.FormItem{
			Text:   "location:",
			Widget: widget.NewLabel("draw A section B number 11"),
		},
		formSpacer,
		&widget.FormItem{
			Text:   "description:",
			Widget: descriptionScroll,
		},
	)

	copyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), nil)
	editButton := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil)
	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
	openButton := widget.NewButton("open", nil)
	actionButtons := container.NewCenter(container.NewHBox(copyButton, spacer, editButton, spacer, deleteButton, spacer, openButton))

	return container.NewVBox(nameBar, spacer, info, spacer, actionButtons)
}

func createFileRow() fyne.CanvasObject {
	file := widget.NewIcon(theme.FileIcon())

	fileName := widget.NewLabel("file name")
	date := widget.NewLabel("01-01-2024")
	space := layout.NewSpacer()
	size := fyne.NewSize(32, 32)
	space.Resize(size)

	row := container.NewHBox(space, file, space, fileName, space, date, space)

	return row
}

func createFilesTable() fyne.CanvasObject {
	pageLangMap := langMap["homePage"][currentLanguage]

	data := make([]string, 100)
	for i := range data {
		data[i] = "file Number " + strconv.Itoa(i)
	}

	// fileInfo := container.NewCenter(container.NewHBox(widget.NewLabel(pageLangMap["noFileSelected"])))
	fileInfo := container.NewCenter(createFileInfoPreview("test file"))

	list := widget.NewList(
		func() int { return len(data) },
		func() fyne.CanvasObject { return createFileRow() },
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[3].(*widget.Label).SetText(data[id])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		fileInfo.RemoveAll()
		fileInfo.Add(createFileInfoPreview(data[id]))
	}

	list.OnUnselected = func(id widget.ListItemID) {
		fileInfo.RemoveAll()
		fileInfo.Add(container.NewHBox(widget.NewLabel(pageLangMap["noFileSelected"])))
	}

	fileTable := container.NewHSplit(list, fileInfo)
	return fileTable
}
