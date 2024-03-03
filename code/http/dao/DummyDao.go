package dao

import (
	"fmt"
	"roomies/code/http/endpoints"
	"roomies/code/shared"
)

type DummyGroceryListDAO struct {
	endpoints.GroceryListDAO
	dummyList *shared.GroceryList
}

func (d *DummyGroceryListDAO) GetList(listID string) *shared.GroceryList {
	return d.dummyList
}

func (d *DummyGroceryListDAO) GetListItem(listID, itemID string) *shared.GroceryListItem {
	return d.dummyList.GetItem(itemID)
}

func (d *DummyGroceryListDAO) CreateList() {
	fmt.Println("CreateList not implemented for dummy")
}
func (d *DummyGroceryListDAO) CreateListItem(listID string, item *shared.GroceryListItem) {
	fmt.Println("CreateListItem not implemented for dummy")
}
func (d *DummyGroceryListDAO) DeleteListItem(listID, itemID string) {
	fmt.Println("Don't delete the dummy list")
}
