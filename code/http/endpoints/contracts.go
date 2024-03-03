package endpoints

import (
	"net/http"
	"roomies/code/shared"
)

type Endpoint interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type GroceryListDAO interface {
	GetList(listID string) *GroceryList
	GetListItem(listID, itemID string) *shared.GroceryListItem
	GetListItems() []*shared.GroceryListItem
	CreateList()
	CreateListItem(listID string, item *shared.GroceryListItem)
	DeleteListItem(listID, itemID string)
}
