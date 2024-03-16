package view

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"roomies/code/app/model"
	"roomies/code/app/util"
)

type RegisterView struct {
	canvasObject   fyne.CanvasObject
	container      *fyne.Container
	mainController *util.MainController
	RegisterModel  model.RegisterModel
}

func NewRegisterView(mainController *util.MainController, registerModel *model.RegisterModel) *RegisterView {
	appView := &RegisterView{
		mainController: mainController,
		RegisterModel:  *registerModel,
	}
	fillerWidget := widget.NewCard("", "", nil)
	registerButton := widget.NewButton("Login", func() { mainController.Window.SetContent(mainController.LoginView.GetCanvasObject()) })
	appView.canvasObject = container.NewBorder(fillerWidget, registerButton, nil, nil, appView.NewRegisterForm())
	return appView
}

func (r *RegisterView) RefreshContent() {}

func (r *RegisterView) GetCanvasObject() fyne.CanvasObject {
	return r.canvasObject
}

func (r *RegisterView) GetViewName() string {
	return "Register"
}

func (r *RegisterView) NewRegisterForm() *widget.Form {
	userNameEntry := widget.NewEntry()
	groupIDEntry := widget.NewEntry()
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
	groupIDEntry.Validator = requiredValidator
	passwordEntry.Validator = requiredValidator

	newForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username: ", Widget: userNameEntry},
			{Text: "Password: ", Widget: passwordEntry},
			{Text: "GroupID: ", Widget: groupIDEntry},
		},
		OnSubmit: func() {
			r.mainController.SwitchToView(r.mainController.AppView)
		},
	}

	return newForm
}
