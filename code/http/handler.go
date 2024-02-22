package main

import (
	"net/http"
)

type Handler struct {
}

func (h Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.RequestURI {
	case "/grocery-list/shared/:groupID":
		if request.Method == "GET" {
			// GET request
		} else {
			// POST request
		}
	case "/grocer-list/shared/:groupID/:itemID":
		// delete an item
	}
}

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/grocery-list/shared/:groupID", Handler{})
	router.Handle("/grocer-list/shared/:groupID/:itemID", Handler{})
	return router
}
