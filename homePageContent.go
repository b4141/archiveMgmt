package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

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

	icon := widget.NewIcon(nil)
	label := widget.NewLabel(pageLangMap["noFileSelected"])
	hbox := container.NewHBox(icon, label)

	list := widget.NewList(
		func() int { return len(data) },
		func() fyne.CanvasObject { return createFileRow() },
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[3].(*widget.Label).SetText(data[id])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(data[id] + ", wow a file")
	}

	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select A File From The List")
	}

	fileTable := container.NewHSplit(list, container.NewCenter(hbox))
	return fileTable
}
