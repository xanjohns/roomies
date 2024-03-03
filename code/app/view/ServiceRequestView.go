package view

import (
	"roomies/code/app/model"
	"roomies/code/app/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ServiceRequestView struct {
	canvasObject        *widget.Form
	mainController      *util.MainController
	ServiceRequestModel model.ServiceRequestModel
}

func NewServiceRequestView(mainController *util.MainController, serviceRequestModel *model.ServiceRequestModel) *ServiceRequestView {
	appView := new(ServiceRequestView)
	appView.mainController = mainController
	appView.ServiceRequestModel = *serviceRequestModel
	appView.canvasObject = appView.NewServiceRequestForm()
	layout.NewFormLayout()
	return appView
}

func (this *ServiceRequestView) RefreshContent() {

}

func (this *ServiceRequestView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *ServiceRequestView) GetViewName() string {
	return "Service Request"
}

func (this *ServiceRequestView) NewServiceRequestForm() *widget.Form {

	problemEntry := widget.NewMultiLineEntry()

	severityOptions := []string{
		"Urgent",
		"Medium",
		"Low",
	}
	severityEntry := widget.NewSelectEntry(severityOptions)
	newForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Describe your problem: ", Widget: problemEntry},
			{Text: "Rate the severity: ", Widget: severityEntry},
		},
		OnSubmit:  func() {

		},
	}

	return newForm
}
