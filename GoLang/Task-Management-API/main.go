package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/govind-yagyasaini/EnchantedForest/Task-Manager-API/login"
	"github.com/govind-yagyasaini/EnchantedForest/Task-Manager-API/registration"
	taskoperations "github.com/govind-yagyasaini/EnchantedForest/Task-Manager-API/task-operations"
)

func main() {
	fmt.Println("Main function is running")
	dbInstance := registration.ConnectPostgresDB()
	defer dbInstance.Close()

	r := mux.NewRouter()
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) { registration.RegisterUser(w, r, dbInstance) }).Methods("POST")
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { login.LoginUser(w, r, dbInstance) }).Methods("POST")
	r.HandleFunc("/task", taskoperations.ProtectedHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":4004", r))

}
