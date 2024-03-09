package view

import (
	"fyne.io/fyne/v2/container"
	"roomies/code/app/model"
	"roomies/code/app/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type RegisterView struct {
	canvasObject   fyne.CanvasObject
	container      *fyne.Container
	mainController *util.MainController
	RegisterModel     model.RegisterModel
}

func NewRegisterView(mainController *util.MainController, loginModel *model.RegisterModel) *RegisterView {
	appView := new(RegisterView)
	appView.mainController = mainController
	appView.RegisterModel = *loginModel
	fillerWidget := widget.NewCard("", "", nil)
	registerButton := widget.NewButton("Login", func() { mainController.Window.SetContent(mainController.LoginView.GetCanvasObject()) })
	appView.canvasObject = container.NewBorder(fillerWidget, registerButton, nil, nil, appView.NewRegisterForm())
	return appView
}

func (this *RegisterView) RefreshContent() {

}

func (this *RegisterView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *RegisterView) GetViewName() string {
	return "Register"
}

func (this *RegisterView) NewRegisterForm() *widget.Form {
	userNameEntry := widget.NewEntry()
	groupIDEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	newForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username: ", Widget: userNameEntry},
			{Text: "Password: ", Widget: passwordEntry},
			{Text: "GroupID: ", Widget: groupIDEntry},
		},
		OnSubmit: func() {
			this.mainController.SwitchToView(this.mainController.AppView)
		},
	}

	return newForm
}
