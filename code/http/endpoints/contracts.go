package endpoints

import "net/http"

type Endpoint interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type GroceryList struct {
}

func (g GroceryList) Handle(w http.ResponseWriter, r *http.Request) {}
