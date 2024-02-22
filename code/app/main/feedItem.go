package main 

import (

	"fyne.io/fyne/v2"
)

type FeedItem fyne.CanvasObject

func NewFeedItem(item fyne.CanvasObject) FeedItem{
	var feedItem FeedItem
	// feedItem.item = item
	// feedItem.time = time.Now()
	return feedItem
}

