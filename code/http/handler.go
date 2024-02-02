package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func (h Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.RequestURI {
	case "/pay-rent":
		fmt.Fprintf(response, "RENT PAID")
	case "/announcement/send":
		fmt.Fprintf(response, "RANDOM ANNOUNCEMENT")
	}
}

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/pay-rent", Handler{})
	router.Handle("/announcement/send", Handler{})
	return router
}
