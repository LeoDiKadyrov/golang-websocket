package registration

import (
	"fmt"
	"net/http"

	registrationValidator "websocket_1/server/registration/validation"
)

// should I connect to db here and do all work here? Or split connection to db into one file and all registration / auth handling to another?

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	http.ServeFile(w, r, "client/registration.html")

	registrationValidator.RegistrationValidator()
}

/* TODO:
- Handle input values on js
- Validate them
- Send POST request to registration.go
- Make http server and handler of POST requests here
- Put DB connection in another file for singleton
- Registration.go does only registration handling (should validation be in another file???)
*/
