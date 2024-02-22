package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GroceryListView struct {
	items []fyne.Widget
	view *fyne.Container

}

func NewGroceryListView() *GroceryListView {
	groceryListView := &GroceryListView{}
	groceryListView.view = container.NewVBox()
	// Get items and add them to view
	groceryListView.view.Add(widget.NewButtonWithIcon("", theme.ContentAddIcon(), groceryListView.AddItem))
	return groceryListView
}


func (this *GroceryListView) AddItem() {
	// dialog.NewForm("New Item Form", "Enter", "Cancel", []*widget.FormItem(widget.NewFormItem("New item:", widget.NewEntry()),), nil, fyne.CurrentApp().Driver().AllWindows()[0]) 
	this.view.Add(widget.NewButtonWithIcon("Paper towels", theme.DeleteIcon(), func() {}))
	
// Will make http request

}

func (this *GroceryListView) GetItems() {

}

func RemoveItem() {

}