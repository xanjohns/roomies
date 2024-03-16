package view

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"roomies/code/app/model"
	"roomies/code/app/util"
)

type LoginView struct {
	canvasObject   fyne.CanvasObject
	container      *fyne.Container
	mainController *util.MainController
	LoginModel     model.LoginModel
}

func NewLoginView(mainController *util.MainController, loginModel *model.LoginModel) *LoginView {
	appView := new(LoginView)
	appView.mainController = mainController
	appView.LoginModel = *loginModel
	fillerWidget := widget.NewCard("", "", nil)
	registerButton := widget.NewButton("Register", func() { mainController.Window.SetContent(mainController.RegisterView.GetCanvasObject()) })
	appView.canvasObject = container.NewBorder(fillerWidget, registerButton, nil, nil, appView.NewLoginForm())
	return appView
}

func (this *LoginView) RefreshContent() {}

func (this *LoginView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *LoginView) GetViewName() string {
	return "Login"
}

func (this *LoginView) NewLoginForm() *widget.Form {
	userNameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	// Create a validation function for required fields
	requiredValidator := func(text string) error {
		if text == "" {
			return errors.New("This field is required")
		}
		return nil
	}

	// Apply validation to entry fields
	userNameEntry.Validator = requiredValidator
	passwordEntry.Validator = requiredValidator

	newForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username: ", Widget: userNameEntry},
			{Text: "Password: ", Widget: passwordEntry},
		},
		OnSubmit: func() {
			this.mainController.SwitchToView(this.mainController.AppView)
		},
	}

	return newForm
}
