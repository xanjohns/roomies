package view

import (
	"roomies/code/app/model"
	"roomies/code/app/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type FeedView struct {
	canvasObject   fyne.CanvasObject
	vBox           fyne.Container
	feedItems      []fyne.CanvasObject
	mainController *util.MainController
	FeedModel      model.FeedModel
}

func NewFeedView(mainController *util.MainController, feedModel *model.FeedModel) *FeedView {
	appView := new(FeedView)
	appView.mainController = mainController
	appView.vBox = *container.NewVBox()
	appView.canvasObject = container.NewVScroll(&appView.vBox)
	appView.FeedModel = *feedModel
	return appView
}

func (this *FeedView) RefreshContent() {
	this.vBox.RemoveAll()
	this.feedItems = nil

	newItems := this.FeedModel.GetItemsHTTP()
	for _, item := range newItems {
		this.AddItem(item)
	}

}

func (this *FeedView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *FeedView) GetViewName() string {
	return "Feed"
}

func (this *FeedView) AddItem(item model.FeedItem) {

}
