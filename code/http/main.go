package main

import (
	"log"
	"net/http"
)

const address = "localhost:8081"

func main() {
	println("Initializing router ... ")
	router := SetupRouter()
	println("Router initialized!")

	println("Starting server ... ")
	log.Fatal(http.ListenAndServe(address, router))
}
