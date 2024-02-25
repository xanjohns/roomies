package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type PeopleView struct {
	canvasObject fyne.CanvasObject
}

func NewPeopleView() *PeopleView {
	appsView := new(PeopleView)
	appsView.canvasObject = widget.NewLabel("PeopleView")
	return appsView
}

func (this *PeopleView) RefreshContent() {

}

func (this *PeopleView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *PeopleView) GetViewName() string {
	return "People"
}
