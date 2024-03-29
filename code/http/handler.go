package main

import (
	"net/http"
	"roomies/code/dao"
	"roomies/code/http/endpoints"
)

type HttpHandler struct {
	endpoint Endpoint
}

func (h HttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	h.endpoint.Handle(response, request)
}

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/grocery-list", HttpHandler{endpoint: endpoints.NewGroceryListHandler(dao.GetNewDummyGroceryListDAO())})
	router.Handle("/announcements", HttpHandler{endpoint: endpoints.NewAnnouncementsHandler(dao.GetNewDummyAnnouncementsDAO())})
	router.Handle("/login", HttpHandler{endpoint: endpoints.NewLoginHandler(dao.GetNewDummyLoginDAO())})
	return router
}
