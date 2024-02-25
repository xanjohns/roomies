package view

import (
	"roomies/code/app/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GroceryView struct {
	canvasObject   *fyne.Container
	groceryItems   map[string]fyne.CanvasObject
	mainController *util.MainController
}

func NewGroceryView(mainController *util.MainController) *GroceryView {
	appsView := new(GroceryView)
	appsView.canvasObject = container.NewVBox()
	appsView.canvasObject.Add(widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {

		// Create a form to get user input
		inputEntry := widget.NewEntry()
		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Enter a string", Widget: inputEntry},
			},
			OnSubmit: func() {
				// Callback function called when the user submits the form
				// Here you can handle the input provided by the user
				println("Entered string:", inputEntry.Text)
				appsView.AddItem(inputEntry.Text)
			},
		}

		// Create a dialog box with the form
		dialog := dialog.NewCustom("String Input Dialog", "OK", form, mainController.Window)

		// Show the dialog box
		dialog.Show()

		// appsView.AddItem("Food")
		appsView.RefreshContent()
	}))
	appsView.groceryItems = map[string]fyne.CanvasObject{}
	return appsView
}

func (this *GroceryView) RefreshContent() {
	// newCanvasObject := container.NewVBox()

}

func (this *GroceryView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *GroceryView) GetViewName() string {
	return "Grocery"
}

func (this *GroceryView) AddItem(item string) {
	if _, exists := this.groceryItems[item]; !exists {

		this.groceryItems[item] = widget.NewButtonWithIcon(item, theme.DeleteIcon(), func() {
			this.canvasObject.Remove(this.groceryItems[item])
			delete(this.groceryItems, item)
			this.canvasObject.Refresh()
		})

		this.canvasObject.Add(this.groceryItems[item])
	}
}
