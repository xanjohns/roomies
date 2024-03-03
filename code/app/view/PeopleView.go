package view

import (
	"roomies/code/app/model"
	"roomies/code/app/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type PeopleView struct {
	canvasObject   fyne.CanvasObject
	vBox           fyne.Container
	contacts       []fyne.CanvasObject
	mainController *util.MainController
	PeopleModel    model.PeopleModel
}

func NewPeopleView(mainController *util.MainController, peopleModel *model.PeopleModel) *PeopleView {
	appView := new(PeopleView)
	appView.mainController = mainController
	appView.vBox = *container.NewVBox()
	appView.canvasObject = container.NewVScroll(&appView.vBox)
	appView.PeopleModel = *peopleModel
	return appView
}

func (this *PeopleView) RefreshContent() {
	this.vBox.RemoveAll()
	this.contacts = nil

	newItems := this.PeopleModel.GetItemsHTTP()
	for _, item := range newItems {
		this.AddItem(item)
	}

}

func (this *PeopleView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *PeopleView) GetViewName() string {
	return "People"
}

func (this *PeopleView) AddItem(item model.Contact) {
	callItem := widget.NewToolbarAction(theme.MediaVideoIcon(), nil)
	textItem := widget.NewToolbarAction(theme.MailSendIcon(), nil)
	EmailItem := widget.NewToolbarAction(theme.MailComposeIcon(), nil)
	infoItem := widget.NewToolbarAction(theme.AccountIcon(), nil)
	toolbar := widget.NewToolbar(infoItem, callItem, textItem, EmailItem)
	border := container.NewBorder(nil, nil, widget.NewLabel(item.Name), toolbar)
	this.vBox.Add(border)
	this.vBox.Add(widget.NewSeparator())
}
