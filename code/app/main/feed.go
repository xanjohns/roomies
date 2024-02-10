package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var countTest int

type Feed struct {
	feedItems *fyne.Container
}

func NewFeed() *Feed {
	myFeed := new(Feed)
	myFeed.feedItems = container.NewVBox()
	countTest = 0
	return myFeed
}

func (this *Feed) AddItem(item fyne.CanvasObject) {
	this.feedItems.Add(item)
}

func (this *Feed) Refresh() {
	log.Println("Refreshing. . .")
	this.getFeed()
}

func (this *Feed) getFeed() {
	countTest++
	this.AddItem(widget.NewLabel(fmt.Sprintf("Post %d", countTest)))
	this.AddItem(widget.NewSeparator())
	// Generate fack feed data
}
