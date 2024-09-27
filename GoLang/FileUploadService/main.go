package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/govindyagyasaini/fileuploadservice/router"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file at the start
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Welcome to File Upload Service in GoLang")
	r := router.Router()

	// Listen and serve on port 4001
	log.Fatal(http.ListenAndServe(":4001", r))
}
