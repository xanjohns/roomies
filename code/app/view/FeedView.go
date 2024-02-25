package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type FeedView struct {
	canvasObject fyne.CanvasObject
}

func NewFeedView() *FeedView {
	appsView := new(FeedView)
	appsView.canvasObject = widget.NewLabel("FeedView")
	return appsView
}

func (this *FeedView) RefreshContent() {

}

func (this *FeedView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *FeedView) GetViewName() string {
	return "Feed"
}
