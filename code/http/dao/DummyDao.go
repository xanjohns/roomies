package dao

import (
	"fmt"
	"roomies/code/http/endpoints"
	"roomies/code/shared"
	"time"
)

type DummyGroceryListDAO struct {
	endpoints.GroceryListDAO
	dummyList *shared.GroceryList
}

func NewDummyGroceryListDAO() *DummyGroceryListDAO {
	return &DummyGroceryListDAO{
		GroceryListDAO: nil,
		dummyList:      nil,
	}
}

func (d *DummyGroceryListDAO) GetList(listID string) *shared.GroceryList {
	return d.dummyList
}

func (d *DummyGroceryListDAO) GetListItems() []*shared.GroceryListItem {
	items := d.dummyList.GetItems()

	if items == nil {
		items = []*shared.GroceryListItem{
			{
				ItemID:        "1111",
				GroupID:       "group1",
				ListItem:      "eggs",
				AddedByID:     "user",
				Recurring:     false,
				RecurringDate: time.Time{},
				Timestamp:     time.Time{},
			}, {
				ItemID:        "2222",
				GroupID:       "group1",
				ListItem:      "milk",
				AddedByID:     "user",
				Recurring:     false,
				RecurringDate: time.Time{},
				Timestamp:     time.Time{},
			}, {
				ItemID:        "3333",
				GroupID:       "group1",
				ListItem:      "flour",
				AddedByID:     "user",
				Recurring:     false,
				RecurringDate: time.Time{},
				Timestamp:     time.Time{},
			}}
	}

	return items
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
