package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"roomies/code/dao"
	"roomies/code/shared"
)

type GroceryListHandler struct {
	responseWriter http.ResponseWriter
	request        *http.Request
	dao            dao.GroceryListDAO
	listID         string
	itemID         string
	listItem       *shared.GroceryListItem
}

func NewGroceryListHandler(listDAO dao.GroceryListDAO) *GroceryListHandler {
	return &GroceryListHandler{
		responseWriter: nil,
		request:        nil,
		dao:            listDAO,
		listID:         "",
		itemID:         "",
		listItem:       nil,
	}
}

func (g *GroceryListHandler) Handle(w http.ResponseWriter, r *http.Request) {
	g.reset()
	g.setup(w, r)

	var res []byte
	var err error

	switch g.request.Method {
	case "GET":
		res, err = json.Marshal(g.getList().GetItems())
		break
	case "POST":
		g.appendItem()
		break
	case "DELETE":
		g.deleteItem()
		break
	}
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	w.Write(res)
}

func (g *GroceryListHandler) appendItem() {
	g.dao.AppendListItem(g.listID, g.listItem)
}
func (g *GroceryListHandler) deleteItem() {
	g.dao.DeleteListItem(g.listID, g.itemID)
}
func (g *GroceryListHandler) getList() *shared.GroceryList {
	return g.dao.GetList(g.listID)
}

func (g *GroceryListHandler) reset() {
	g.request = nil
	g.responseWriter = nil
	g.listID = ""
	g.itemID = ""
}

func (g *GroceryListHandler) setup(w http.ResponseWriter, r *http.Request) {
	g.responseWriter = w
	g.request = r
	if r.Body != nil {
		json.NewDecoder(r.Body).Decode(g.listItem)
	}
	g.listID = r.URL.Query().Get("listId")
	g.itemID = r.URL.Query().Get("itemId")
}
