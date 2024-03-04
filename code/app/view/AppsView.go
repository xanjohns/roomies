package view

import (
	"roomies/code/app/model"
	"roomies/code/app/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppsView struct {
	canvasObject   *fyne.Container
	mainController *util.MainController
	appViews       []util.RoomiesView
}

func NewAppsView(mainController *util.MainController) *AppsView {
	appsView := new(AppsView)
	appsView.canvasObject = container.NewVBox()
	appsView.AddApp(NewGroceryView(mainController, model.NewGroceryModel()))
	appsView.AddApp(NewAnnouncementsView(mainController, model.NewAnnouncementsModel()))
	appsView.AddApp(NewServiceRequestView(mainController, model.NewServiceRequestModel()))
	appsView.mainController = mainController
	return appsView
}

func (this *AppsView) RefreshContent() {
	// this.CreateAppsCanvasObject()
}

func (this *AppsView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *AppsView) GetViewName() string {
	return "Apps"
}

func (this *AppsView) AddApp(newApp util.RoomiesView) {
	this.appViews = append(this.appViews, newApp)
	this.canvasObject.Add(widget.NewButton(newApp.GetViewName(), func() {this.mainController.SwitchToView((newApp))}))
}
