package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/govindyagyasaini/fileuploadservice/router"
)

func main() {

	fmt.Println("Welcome to File Upload Service in GoLang")
	router := router.Router()
	log.Fatal(http.ListenAndServe(":4001", router))

}
