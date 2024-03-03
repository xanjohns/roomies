package endpoints

import (
	"encoding/json"
	"fmt"
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

func NewGroceryList() *GroceryList {
	return &GroceryList{
		responseWriter: nil,
		request:        nil,
		dao:            nil,
		listType:       "",
		listID:         "",
		itemID:         "",
		listItem:       nil,
	}
}

func (g *GroceryList) Handle(w http.ResponseWriter, r *http.Request) {
	g.setup(w, r)
	//defer g.reset()

	switch g.request.Method {
	case "GET":
		//g.List()
		items := []*shared.GroceryListItem{
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

		res, err := json.Marshal(items)
		if err != nil {
			fmt.Fprintf(w, "errrrrrror")
		}
		//fmt.Fprintf(w, "{%s, %s, %s}", items[0].ListItem, items[1].ListItem, items[2].ListItem)
		w.Write(res)
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
	g.dao.GetListItems()

}
func (g *GroceryList) GetAll() {
	return
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
	g.listType = splitURI[2]
	g.listID = splitURI[3]
	//if len(splitURI) > 3 {
	//	g.itemID = splitURI[4]
	//}
}
