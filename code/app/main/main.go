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

	tabs := container.NewAppTabs(
		container.NewTabItem("Feed", container.NewVScroll(
			feed.feedItems)),

		// container.NewTabItem("Apps", scrollPadded),
		container.NewTabItem("People", widget.NewLabel("Peoples")),
	)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			feed.Refresh()
		}),
	)

	tabs.SetTabLocation(container.TabLocationBottom)
	tabs.Resize(fyne.NewSize(60, 300))
	content := container.NewBorder(toolbar, nil, nil, nil, tabs)

	w.SetContent(content)
	w.ShowAndRun()

}
