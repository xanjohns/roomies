package view

import (
	"roomies/code/app/model"
	"roomies/code/app/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AnnouncementsView struct {
	canvasObject       *fyne.Container
	announcements      []fyne.CanvasObject
	mainController     *util.MainController
	AnnouncementsModel model.AnnouncementsModel
}

func NewAnnouncementsView(mainController *util.MainController, announcementsModel *model.AnnouncementsModel) *AnnouncementsView {
	appView := new(AnnouncementsView)
	appView.mainController = mainController
	appView.AnnouncementsModel = *announcementsModel
	appView.canvasObject = container.NewVBox()
	return appView
}

func (this *AnnouncementsView) RefreshContent() {
	this.canvasObject.RemoveAll()

	this.announcements = nil

	newItems := this.AnnouncementsModel.GetItemsHTTP()
	for _, item := range newItems {
		this.AddItem(item)
	}
}

func (this *AnnouncementsView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *AnnouncementsView) GetViewName() string {
	return "Announcements"
}

func (this *AnnouncementsView) AddItem(item model.Announcement) {
	newCard := widget.NewCard(item.Message, (item.Date + " -- " + item.From), nil)

	this.canvasObject.Add(newCard)
}
