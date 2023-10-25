package authentication

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"errors"
	"log"

	"websocket_1/server/models"
	"websocket_1/server/validation"
	"golang.org/x/crypto/bcrypt"
	postgresdb "websocket_1/server/database"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var request models.User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := request.Username
	password := request.PasswordHash

	if !validation.IsValidPassword(password) || !validation.IsValidUsername(username) {
		http.Error(w, "Invalid authentication input", http.StatusBadRequest)
		return
	}

	if err := authenticateUser(username, password); err != nil {
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Authentication successful"))
}

func authenticateUser(username, password string) error {
	var user models.User
	db := postgresdb.GetInstanceDB().DB

	stmt, err := db.Prepare("SELECT username, password_hash FROM users WHERE username = $1;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid username")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		log.Fatal(err)
	} // IF ERROR SHOW ON FRONTEND

	return nil
}
