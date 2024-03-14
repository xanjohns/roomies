package dao

import "roomies/code/shared"

type GroceryListDAO interface {
	GetList(listID string) *shared.GroceryList
	GetListItem(listID, itemID string) *shared.GroceryListItem
	CreateList()
	AppendListItem(listID string, item *shared.GroceryListItem)
	DeleteListItem(listID, itemID string)
}

type LoginDAO interface {
	ValidateLogin(username, password string) bool
}
