package view

import (
	"roomies/code/app/model"
	"roomies/code/app/util"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type LoginView struct {
	canvasObject *widget.Form
	mainController *util.MainController
	LoginModel model.LoginModel
}

func NewLoginView(mainController *util.MainController, loginModel *model.LoginModel) *LoginView {
	appView := new(LoginView)
	appView.mainController = mainController
	appView.LoginModel = *loginModel
	appView.canvasObject = appView.NewLoginForm()
	layout.NewFormLayout()
	return appView
}

func (this *LoginView) RefreshContent() {

}

func (this *LoginView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *LoginView) GetViewName() string {
	return "Login"
}

func (this *LoginView) NewLoginForm() *widget.Form {
	userNameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	newForm := &widget.Form {
		Items: []*widget.FormItem {
			{Text: "Username: ", Widget: userNameEntry},
			{Text: "Password: ", Widget: passwordEntry},
		},
		OnSubmit: func() {
			this.mainController.SwitchToView(this.mainController.AppView)
		},
	}

	return newForm
}