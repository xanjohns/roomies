package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const appsButtonSizeHeight = 100
const appsButtonSizeWidth = 100

func main() {
	a := app.New()
	w := a.NewWindow("Roomies")
	w.Resize(fyne.NewSize(600, 600))

	feed := NewFeed()

	// tabs := container.NewAppTabs(
	// 	container.NewTabItem("Feed", container.NewVScroll(
	// 		feed.feedItems)),

	// 	// container.NewTabItem("Apps", scrollPadded),
	// 	container.NewTabItem("People", widget.NewLabel("Peoples")),
	// )

	mainView := container.NewVBox()

	appsView := NewAppsView(w)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			feed.Refresh()
		}),
	)

	tabs := container.NewHBox(
		widget.NewButton("Feed", func() {}),
		widget.NewButton("Apps", func() {
			mainView.Objects = appsView.apps
			mainView.Refresh()
		}),
		widget.NewButton("People", func() {}),
	)

	centerTabs := container.NewCenter(tabs)

	content := container.NewBorder(toolbar, centerTabs, nil, nil, mainView)

	w.SetContent(content)
	w.ShowAndRun()

}

func setView(w fyne.Window, top fyne.CanvasObject, bottom fyne.CanvasObject, mid fyne.CanvasObject) {
	content := container.NewBorder(top, bottom, nil, nil, mid)
	w.SetContent(content)
	w.Show()

}
