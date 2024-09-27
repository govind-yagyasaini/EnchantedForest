package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/govind-yagyasaini/EnchantedForest/GoLang/GoWeatherAPI/controller"
)

func Router(apiKey string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/location/{city}", func(w http.ResponseWriter, r *http.Request) {
		GetCurrentWeather(w, r, apiKey)
	}).Methods("GET")
	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("<h1>Welcome to Go Weather API.</h1>")
}

func GetCurrentWeather(w http.ResponseWriter, r *http.Request, apiKey string) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	weather, err := controller.CurrentByName(params["city"], apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(weather)
	if err != nil {
		log.Fatal(err)
	}
}
