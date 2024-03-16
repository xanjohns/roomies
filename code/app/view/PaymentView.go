package view

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"roomies/code/app/model"
	"roomies/code/app/util"
)

type PaymentView struct {
	canvasObject   fyne.CanvasObject
	container      *fyne.Container
	mainController *util.MainController
	PaymentModel   model.PaymentModel
}

func NewPaymentView(mainController *util.MainController, PaymentModel *model.PaymentModel) *PaymentView {
	appView := &PaymentView{
		mainController: mainController,
		PaymentModel:   *PaymentModel,
	}

	// Create a label at the top of the view with custom theme
	titleLabel := widget.NewLabelWithStyle("Payment", fyne.TextAlignCenter, getCustomTextStyle())

	// Set the canvasObject with a Border container
	appView.canvasObject = container.NewBorder(titleLabel, nil, nil, nil, appView.NewPaymentForm())

	return appView
}

func getCustomTextStyle() fyne.TextStyle {
	return fyne.TextStyle{Bold: true, Italic: true, Monospace: false}
}

func (this *PaymentView) RefreshContent() {
	// Implement content refresh logic here if needed
}

func (this *PaymentView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *PaymentView) GetViewName() string {
	return "Payment"
}

func (this *PaymentView) NewPaymentForm() fyne.CanvasObject {
	// Existing widget declarations...
	FirstNameEntry := widget.NewEntry()
	LastNameEntry := widget.NewEntry()
	addressEntry := widget.NewEntry()
	cityEntry := widget.NewEntry()
	stateEntry := widget.NewEntry()
	creditCardEntry := widget.NewEntry()
	cvcEntry := widget.NewEntry()

	// Create a dropdown for selecting month with custom theme
	monthSelect := widget.NewSelect([]string{
		"January", "February", "March", "April", "May", "June", "July",
		"August", "September", "October", "November", "December",
	}, func(selected string) {
		// Handle month selection
	})

	// Create a dropdown for selecting year with custom theme
	yearSelect := widget.NewSelect([]string{
		"2024", "2025", "2026", "2027", "2028", // Add more years as needed
	}, func(selected string) {
		// Handle year selection
	})

	// Create validation functions for required fields
	requiredValidator := func(text string) error {
		if text == "" {
			return errors.New("This field is required")
		}
		return nil
	}

	// Apply validation to entry fields
	FirstNameEntry.Validator = requiredValidator
	LastNameEntry.Validator = requiredValidator
	addressEntry.Validator = requiredValidator
	cityEntry.Validator = requiredValidator
	stateEntry.Validator = requiredValidator
	creditCardEntry.Validator = requiredValidator
	cvcEntry.Validator = requiredValidator

	// Create a new Form with labeled entry fields and the expiration date dropdowns
	newForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "First Name:", Widget: FirstNameEntry},
			{Text: "Last Name:", Widget: LastNameEntry},
			{Text: "Address:", Widget: addressEntry},
			{Text: "City:", Widget: cityEntry},
			{Text: "State:", Widget: stateEntry},
			{Text: "Credit Card Number:", Widget: creditCardEntry},
			{Text: "CVC:", Widget: cvcEntry},
			{Text: "Expiration Date:", Widget: container.NewHBox(monthSelect, yearSelect)},
		},
		OnSubmit: func() {
			this.mainController.SwitchToView(this.mainController.AppView)
		},
	}

	return newForm
}
