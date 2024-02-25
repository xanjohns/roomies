package endpoints

import (
	"net/http"
	"strings"
)

type GroceryList struct {
	responseWriter http.ResponseWriter
	request        *http.Request
	listType       string
	groupID        string
	itemID         string
}

func (g *GroceryList) Handle(w http.ResponseWriter, r *http.Request) {
	g.setup(w, r)
	defer g.reset()

	switch g.listType {
	case "shared":
		g.sharedListHandler()
		break
	case "individual":
		g.individualListHandler()
		break
	}
}

func (g *GroceryList) sharedListHandler() {
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
func (g *GroceryList) individualListHandler() {}

func (g *GroceryList) Add()      {}
func (g *GroceryList) Delete()   {}
func (g *GroceryList) List()     {}
func (g *GroceryList) Retrieve() {}

func (g *GroceryList) reset() {
	g.request = nil
	g.responseWriter = nil
	g.listType = ""
	g.groupID = ""
	g.itemID = ""
}

func (g *GroceryList) setup(w http.ResponseWriter, r *http.Request) {
	g.responseWriter = w
	g.request = r
	splitURI := strings.Split(r.RequestURI, "/")
	g.listType = splitURI[1]
	g.groupID = splitURI[2]
	if len(splitURI) > 3 {
		g.itemID = splitURI[3]
	}
}
