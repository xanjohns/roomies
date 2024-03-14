package endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"roomies/code/dao"
	"roomies/code/shared"
	"testing"
	"time"
)

func TestHandleGET(t *testing.T) {
	handler := NewGroceryListHandler(dao.GetNewDummyGroceryListDAO())
	req, err := http.NewRequest("GET", "/?listId=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler.Handle(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var items []*shared.GroceryListItem
	err = json.Unmarshal(w.Body.Bytes(), &items)
	if err != nil {
		t.Fatal(err)
	}

	testItems := dao.GetNewDummyGroceryListDAO().GetList("x").GetItems()

	if len(items) == 0 {
		t.Errorf("No Items Found")
	}

	for i := range items {
		if items[i].ListItem != testItems[i].ListItem {
			t.Errorf("Got item %s but expected item %s", items[i].ListItem, testItems[i].ListItem)
		}
	}
}

func TestHandlePOST(t *testing.T) {
	handler := NewGroceryListHandler(dao.GetNewDummyGroceryListDAO())
	item := shared.GetNewGroceryListItem("xx", "listitem", false, time.Time{})

	jsonItem, err := json.Marshal(item)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/?listId=123", bytes.NewBuffer(jsonItem))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler.Handle(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Add your assertions based on expected response
}

func TestHandleDELETE(t *testing.T) {
	handler := NewGroceryListHandler(dao.GetNewDummyGroceryListDAO())
	req, err := http.NewRequest("DELETE", "/?listId=123&itemId=456", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler.Handle(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Add your assertions based on expected response
}
