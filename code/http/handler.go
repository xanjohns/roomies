package main

import (
	"net/http"
	"roomies/code/http/endpoints"
)

type HttpHandler struct {
	endpoint endpoints.Endpoint
}

func (h HttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	h.endpoint.Handle(response, request)
}

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	//list := endpoints.NewGroceryList()
	router.Handle("/grocery-list", HttpHandler{endpoint: &endpoints.GroceryList{}})
	router.Handle("/grocery-list/shared/group1", HttpHandler{endpoint: &endpoints.GroceryList{}})
	router.Handle("/grocery-list/private", HttpHandler{endpoint: &endpoints.GroceryList{}})

	return router
}
