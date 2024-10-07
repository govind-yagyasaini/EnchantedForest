package taskoperations

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/govind-yagyasaini/EnchantedForest/Task-Manager-API/login"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Malformed authorization header")
		return
	}

	tokenString := headerParts[1] // This should be the correct token string
	fmt.Println("Extracted token string:", tokenString)

	err := login.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	fmt.Fprint(w, "Welcome to the protected area")

	//Enter all your task CRUD operations here

}
