package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/govind-yagyasaini/EnchantedForest/GoLang/GoWeatherAPI/router"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting Go Weather API.")

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Fetch the API key
	apiKey := os.Getenv("OWM_API_KEY")
	if apiKey == "" {
		log.Fatal("OWM_API_KEY is not set")
	}

	// Pass the API key to the router
	r := router.Router(apiKey)

	log.Fatal(http.ListenAndServe(":4002", r))
}
