package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/govind-yagyasaini/EnchantedForest/Go-URL-Shortner/awssession"
	"github.com/govind-yagyasaini/EnchantedForest/Go-URL-Shortner/handlers"
)

func main() {
	fmt.Println("Welcome to URL Shortener Service in GoLang")

	// Initialize DynamoDB Client
	awssession.InitDynamoDBClient()

	// Define URL routes
	http.HandleFunc("/shorten", handlers.HandleShorten)
	http.HandleFunc("/short/", handlers.HandleRedirect)

	fmt.Println("URL Shortener is running on :4003")
	log.Fatal(http.ListenAndServe(":4003", nil)) // Use log.Fatal to catch errors
}
