package main

import (
	"log"
	"net/http"
)

func main() {
	println("initializing router ... ")
	router := SetupRouter()

	println("starting server ... ")
	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
