package view

import (
	"roomies/code/app/model"
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
	GroceryModel   model.GroceryModel
}

func NewGroceryView(mainController *util.MainController, groceryModel *model.GroceryModel) *GroceryView {
	appView := new(GroceryView)
	appView.mainController = mainController
	appView.GroceryModel = *groceryModel
	appView.canvasObject = container.NewVBox()
	appView.canvasObject.Add(appView.NewAddButton())
	appView.groceryItems = map[string]fyne.CanvasObject{}
	return appView
}

func (this *GroceryView) RefreshContent() {
	this.canvasObject.RemoveAll()
	for k := range this.groceryItems {
		delete(this.groceryItems, k)
	}

	this.canvasObject.Add(this.NewAddButton())
	newItems := this.GroceryModel.GetItemsHTTP()
	for _, item := range newItems {
		this.AddItem(item)
	}

}

func (this *GroceryView) GetCanvasObject() fyne.CanvasObject {
	return this.canvasObject
}

func (this *GroceryView) GetViewName() string {
	return "Grocery"
}

func (this *GroceryView) AddItem(item model.GroceryListItem) {
	if _, exists := this.groceryItems[item.ListItem]; !exists {

		this.groceryItems[item.ListItem] = widget.NewButtonWithIcon(item.ListItem, theme.DeleteIcon(), func() {
			this.canvasObject.Remove(this.groceryItems[item.ListItem])
			delete(this.groceryItems, item.ListItem)
			this.canvasObject.Refresh()
		})

		this.canvasObject.Add(this.groceryItems[item.ListItem])
	}
}

func (this *GroceryView) NewAddButton() fyne.CanvasObject {
	newWidget := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {

		// Create a form to get user input
		inputEntry := widget.NewEntry()
		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Enter an item", Widget: inputEntry},
			},
			OnSubmit: func() {
				// Callback function called when the user submits the form
				// Here you can handle the input provided by the user
				println("Entered item:", inputEntry.Text)
				this.AddItem(model.GroceryListItem{ListItem: inputEntry.Text})
			},
		}

		// Create a dialog box with the form
		dialog := dialog.NewCustom("Item Input Dialog", "OK", form, this.mainController.Window)

		// Show the dialog box
		dialog.Show()

		// appsView.AddItem("Food")
		this.RefreshContent()
	})
	return newWidget
}
