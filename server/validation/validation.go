package validation

import (
	"regexp"
	"unicode"
)

func IsValidPassword(password string) bool {
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

func IsValidUsername(username string) bool {
	var englishLettersPattern = regexp.MustCompile("^[a-zA-Z]+$")
	var usernameLength = len(username)

	if usernameLength >= 4 && usernameLength <= 40 && englishLettersPattern.MatchString(username) {
		return true
	} else {
		return false
	}
}