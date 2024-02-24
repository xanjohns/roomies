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
	router.Handle("/grocery-list", HttpHandler{endpoint: endpoints.GroceryList{}})
	return router
}
