package login

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func VerifyHashedPassword(hashedpassword string, password string) error {
	fmt.Printf("Stored hashed password: %s\n", hashedpassword)
	fmt.Printf("Incoming plaintext password: %s\n", password)

	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	if err != nil {
		fmt.Println("Password does not match:", err)
	} else {
		fmt.Println("Password matches!")
	}
	return err
}

func LoginUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	var credentials *Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Message body is incorrect.", http.StatusBadRequest)
		return
	}
	var storedHashedPassword string
	sqlStatement := `Select password from registration where username = $1`
	err = db.QueryRow(sqlStatement, credentials.Username).Scan(&storedHashedPassword)
	if err != nil {
		http.Error(w, "Username or Password is incorrect.", http.StatusBadRequest)
		return
	}
	fmt.Println("Stored hashed password from DB:", storedHashedPassword)
	fmt.Println("Plaintext password from user input:", credentials.Password)

	err = VerifyHashedPassword(storedHashedPassword, credentials.Password)
	if err != nil {
		http.Error(w, "Password is incorrect.", http.StatusBadRequest)
		return
	}

	jwtToken, err := CreateToken(credentials.Username)
	if err != nil {
		http.Error(w, "Error while creating JWT token.", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"jwtToken": jwtToken})

}
