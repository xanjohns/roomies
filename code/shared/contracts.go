package shared

import (
	"time"
)

type GroceryList struct {
	items []*GroceryListItem
}

func (g *GroceryList) GetItem(itemID string) *GroceryListItem {
	for _, item := range g.items {
		if item.ItemID == itemID {
			return item
		}
	}
	return nil
}

func (g *GroceryList) GetItems() []*GroceryListItem {
	return g.items
}

func GetNewGroceryList(items []*GroceryListItem) *GroceryList {
	if items == nil {
		items = make([]*GroceryListItem, 0)
	}
	return &GroceryList{items: items}
}

type GroceryListItem struct {
	ListID        string
	ItemID        string
	GroupID       string
	ListItem      string
	AddedByID     string
	Recurring     bool
	RecurringDate time.Time
	Timestamp     time.Time
}

func GetNewGroceryListItem(groupID, listItem string, recurring bool, recurringDate time.Time) *GroceryListItem {
	gli := GroceryListItem{
		ItemID:        "", //TODO does SQL generate this? Or do we need to generate this uniquely?
		GroupID:       groupID,
		ListItem:      listItem,
		AddedByID:     "The user", //TODO probably have a central object to get this information from
		Recurring:     recurring,
		RecurringDate: recurringDate,
		Timestamp:     time.Now(),
	}
	return &gli
}
