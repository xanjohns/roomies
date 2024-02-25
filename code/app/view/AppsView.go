package view

import (
	"roomies/code/app/model"
	"roomies/code/app/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppsView struct {
	canvasObject   fyne.CanvasObject
	mainController *util.MainController
	appViews       []util.RoomiesView
}

func NewAppsView(mainController *util.MainController) *AppsView {
	appsView := new(AppsView)
	appsView.AddApp(NewGroceryView(mainController, model.NewGroceryModel()))
	appsView.mainController = mainController
	appsView.canvasObject = widget.NewLabel("AppsView")
	return appsView
}

func (this *AppsView) RefreshContent() {
	this.CreateAppsCanvasObject()
}

func (this *AppsView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *AppsView) CreateAppsCanvasObject() {
	canvasObject := container.NewVBox()

	for _, app := range this.appViews {
		canvasObject.Add(widget.NewButton(app.GetViewName(), func() { this.mainController.SwitchToView(app) }))
	}

	this.canvasObject = canvasObject
}

func (this *AppsView) GetViewName() string {
	return "Apps"
}

func (this *AppsView) AddApp(newApp util.RoomiesView) {
	this.appViews = append(this.appViews, newApp)
}
