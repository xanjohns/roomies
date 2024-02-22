package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
)

func NewMainView(win fyne.Window, middle fyne.CanvasObject) *fyne.Container {
	mainView := container.NewVBox()

	appsView := NewAppsView(win)

	feed := NewFeed()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			feed.Refresh()
		}),
	)

	tabs := container.NewHBox(
		widget.NewButton("Feed", func() {}),
		widget.NewButton("Apps", func() {
			mainView.Objects = appsView.apps
			win.SetContent(NewMainView(win, mainView))
		}),
		widget.NewButton("People", func() {}),
	)


	centerTabs := container.NewCenter(tabs)

	content := container.NewBorder(toolbar, centerTabs, nil, nil, middle)

	return content

}