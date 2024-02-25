package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GroceryView struct {
	canvasObject *fyne.Container
	groceryItems map[string]fyne.CanvasObject
}

func NewGroceryView() *GroceryView {
	appsView := new(GroceryView)
	appsView.canvasObject = container.NewVBox()
	appsView.canvasObject.Add(widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		appsView.AddItem("Food")
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
