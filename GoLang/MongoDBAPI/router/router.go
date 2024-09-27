package router

import (
	"github.com/gorilla/mux"
	"github.com/govindyagyasaini/mongodbapi/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api", controller.HomeHandler).Methods("GET")
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.InsertMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")
	return r

}
