package model

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type GroceryModel struct {
}

type GroceryListItem struct {
	ItemID   string
	GroupID  string
	ListItem string
	AddedByID string
	Recurring bool
	RecurringDate time.Time
	Timestamp time.Time
}

func NewGroceryModel() *GroceryModel {
	return new(GroceryModel)
}

func (this *GroceryModel) PostItemHTTP() {

}

func (this *GroceryModel) GetItemsHTTP() []GroceryListItem {
	response, err := http.Get("http://localhost:8081/grocery-list/shared/group1")
	if err != nil {
		log.Println(err)
		return nil
	}
	data, _ := io.ReadAll(response.Body)

	var items []GroceryListItem
	if err := json.Unmarshal(data, &items); err != nil {
		log.Println(err)
	}
	log.Println(items)

	// Dummy data

	// items := []string {
	// 	"Eggs",
	// 	"Milk",
	// 	"Flour",
	// }

	return items
}
