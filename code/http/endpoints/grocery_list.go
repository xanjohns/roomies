package endpoints

import (
	"net/http"
	"roomies/code/shared"
	"strings"
	"time"
)

type GroceryList struct {
	responseWriter http.ResponseWriter
	request        *http.Request
	dao            GroceryListDAO
	listType       string //TODO currently there doesn't seem a need to have a different handling of "shared" and "individual" list types
	listID         string
	itemID         string
	listItem       *shared.GroceryListItem
}

func (g *GroceryList) Handle(w http.ResponseWriter, r *http.Request) {
	g.setup(w, r)
	defer g.reset()

	switch g.request.Method {
	case "GET":
		g.List()
		break
	case "POST":
		g.Add()
		break
	case "DELETE":
		g.Delete()
		break
	}
}

func (g *GroceryList) Add() {
	g.dao.CreateListItem(g.listID, g.listItem)
}
func (g *GroceryList) Delete() {
	g.dao.DeleteListItem(g.listID, g.itemID)
}
func (g *GroceryList) List() {
	g.dao.GetList(g.listID)
}
func (g *GroceryList) Retrieve() {
	g.dao.GetListItem(g.listID, g.itemID)
}

func (g *GroceryList) reset() {
	g.request = nil
	g.responseWriter = nil
	g.listType = ""
	g.listID = ""
	g.itemID = ""
}

func (g *GroceryList) setup(w http.ResponseWriter, r *http.Request) {
	g.responseWriter = w
	g.request = r
	g.listItem = shared.GetNewGroceryListItem("", "", false, time.Time{})
	splitURI := strings.Split(r.RequestURI, "/")
	g.listType = splitURI[1]
	g.listID = splitURI[2]
	if len(splitURI) > 3 {
		g.itemID = splitURI[3]
	}
}
