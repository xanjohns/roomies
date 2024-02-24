package endpoints

import "net/http"

type GroceryList struct {
}

func (g *GroceryList) Handle(w http.ResponseWriter, r *http.Request) {

}

func (g *GroceryList) Add()      {}
func (g *GroceryList) Delete()   {}
func (g *GroceryList) List()     {}
func (g *GroceryList) Retrieve() {}
