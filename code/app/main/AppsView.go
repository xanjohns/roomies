package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type AppsView struct {
	apps []fyne.CanvasObject
	w fyne.Window
}

func NewAppsView(win fyne.Window) *AppsView {
	return &AppsView{
		apps: []fyne.CanvasObject{
			widget.NewButton("Shopping List", func() {
				win.SetContent(NewMainView(win, NewGroceryListView().view))
			}),
			widget.NewButton("Pay rent", func() {}),
			widget.NewButton("Schedule", func() {}),
			widget.NewButton("Announcements", func() {}),
		},
		w: win,
	}
}