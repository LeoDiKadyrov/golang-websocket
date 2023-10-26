package customError

import "net/http"

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}

func SendCustomHttpError(w http.ResponseWriter, statusCode int, message string) {
	customError := &CustomError{
		StatusCode: statusCode,
		Message: message,
	}
	http.Error(w, customError.Error(), statusCode)
}