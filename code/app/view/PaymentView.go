package view

import (
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

	// Create a label at the top of the view
	titleLabel := widget.NewLabelWithStyle("Payment", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: true})

	// Create a filler widget (replace nil with actual content)
	//fillerWidget := widget.NewCard("", "", nil)

	// Set the canvasObject with a Border container
	appView.canvasObject = container.NewBorder(titleLabel, nil, nil, nil, appView.NewPaymentForm())

	return appView
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

func (this *PaymentView) NewPaymentForm() *widget.Form {
	FirstNameEntry := widget.NewEntry()
	LastNameEntry := widget.NewEntry()
	addressEntry := widget.NewEntry()
	cityEntry := widget.NewEntry()
	stateEntry := widget.NewEntry()
	creditCardEntry := widget.NewEntry()
	cvcEntry := widget.NewEntry()

	// Create a new Form with labeled entry fields
	newForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "First Name:", Widget: FirstNameEntry},
			{Text: "Last Name:", Widget: LastNameEntry},
			{Text: "Address:", Widget: addressEntry},
			{Text: "City:", Widget: cityEntry},
			{Text: "State:", Widget: stateEntry},
			{Text: "Credit Card Number:", Widget: creditCardEntry},
			{Text: "CVC:", Widget: cvcEntry},
		},
		OnSubmit: func() {
			this.mainController.SwitchToView(this.mainController.AppView)
		},
	}

	return newForm
}
