package registration

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Import PostgreSQL driver
	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "username"
	password = "password"
	dbname   = "userregistration"
)

var connectionURL = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func ConnectPostgresDB() *sql.DB {
	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		log.Printf("Error occurred while creating DB connection - %s\n", err)

	} else {
		fmt.Println("Connection successful.")
	}

	return db

}

func InsertIntoPostgresDB(db *sql.DB, username, email, password string) {
	sqlStatement := `INSERT INTO registration (username, email, password) 
    VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, username, email, password)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println(err)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return
	}

	// Call InsertIntoPostgresDB with the user's data
	InsertIntoPostgresDB(db, user.Username, user.Email, string(bytes))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"user": "User registered successfully."})

}
