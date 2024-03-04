package model

import (
	// "encoding/json"
	// "io"
	// "net/http"
	// "log"
)

type GroceryModel struct {
}

func NewGroceryModel() *GroceryModel {
	return new(GroceryModel)
}

func (this *GroceryModel) PostItemHTTP() {

}

func (this *GroceryModel) GetItemsHTTP() []string {
	// response, err := http.Get("http://localhost:8080/grocery-list/shared/{id}")
	// if err != nil {
	// 	log.Println(err)
	// 	return nil
	// }
	// data, _ := io.ReadAll(response.Body)

	// var items []string
	// if err := json.Unmarshal(data, &items); err != nil {
	// 	log.Println(err)
	// }

	items := []string {
		"Eggs",
		"Milk",
		"Flour",
	}

	return items
}
