package shared

import "time"

type GroceryList struct {
	items map[string]*GroceryListItem
}

func (g *GroceryList) GetItem(itemID string) *GroceryListItem {
	return g.items[itemID]
}

func GetNewGroceryList() *GroceryList {
	return &GroceryList{items: make(map[string]*GroceryListItem)}
}

type GroceryListItem struct {
	ItemID        int
	GroupID       string
	ListItem      string
	AddedByID     string
	Recurring     bool
	RecurringDate time.Time
	Timestamp     time.Time
}

func GetNewGroceryListItem(groupID, listItem string, recurring bool, recurringDate time.Time) *GroceryListItem {
	gli := GroceryListItem{
		ItemID:        0, //TODO does SQL generate this? Or do we need to generate this uniquely?
		GroupID:       groupID,
		ListItem:      listItem,
		AddedByID:     "The user", //TODO probably have a central object to get this information from
		Recurring:     recurring,
		RecurringDate: recurringDate,
		Timestamp:     time.Now(),
	}
	return &gli
}
