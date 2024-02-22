package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Apps struct {
	appItems *fyne.Container
}

func NewApps() *Apps {
	myApps := new(Apps)
	myApps.appItems = container.NewVBox()

	// Default apps
	// myApps.AddItem(widget.NewButton(""))

	return myApps
}

func (this *Apps) AddItem(item fyne.CanvasObject) {
	this.appItems.Add(item)
	this.appItems.Add(layout.NewSpacer())
}
