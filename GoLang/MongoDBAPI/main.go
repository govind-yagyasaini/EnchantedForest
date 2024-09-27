package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/govindyagyasaini/mongodbapi/router"
)

func main() {

	fmt.Println("Welcome to GoLang API with MongoDB")
	router := router.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
}
