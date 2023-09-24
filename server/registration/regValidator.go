package registration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"unicode"
)

type UserRegistration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegValidator(w http.ResponseWriter, r *http.Request) {
	var request UserRegistration
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := request.Username
	password := request.Password

	fmt.Println("username in validator.go: ", username)
	fmt.Println("password in validator.go: ", password)

	if !isValidPassword(password) || !isValidUsername(username) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Perform registration
	if err := RegisterUser(username, password); err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	// Registration successful // TODO: PUT THIS INTO REGISTRATION.GO
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration successful"))
}

func isValidPassword(password string) bool {
	letters := 0
	var number, upper, special, eightOrMore bool
	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			number = true
		case unicode.IsUpper(char):
			upper = true
			letters++
		case unicode.IsLetter(char) || char == ' ':
			letters++
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			special = true
		default:
			return false
		}
	}
	eightOrMore = letters >= 8
	if number && upper && special && eightOrMore {
		return true
	}

	return false
}

func isValidUsername(username string) bool {
	var englishLettersPattern = regexp.MustCompile("^[a-zA-Z]+$")
	var usernameLength = len(username)

	if usernameLength >= 4 && usernameLength <= 40 && englishLettersPattern.MatchString(username) {
		return true
	} else {
		return false
	}
}
