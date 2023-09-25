package authentication

import (
	"fmt"
)

func Authentication() {
	fmt.Println("authentication package")
	// accept validated data from frontend
	// check in database if users exists
	// if exists - approve and redirect with successfull message
	// if not - disapprove and show red modal
}