package helper

import "net/http"

const (
	// User Input
	ErrorUserInput       = "the data sent is incorrect"
	ErrorUserInputFormat = "data format is not supported"
	ErrorUserCredential  = "incorrect email or password"
	ErrorInvalidValidate = "validation is invalid"

	// Server
	ErrorGeneralServer = "an error occurred in the server process"

	// Database
	ErrorGeneralDatabase  = "there is a problem with the database"
	ErrorNoRowsAffected   = "no changes to the database"
	ErrorDatabaseNotFound = "no data found on the database"
)

func ErrorCode(e error) int {
	if e == nil {
		return 200
	}
	switch e.Error() {
	// User Input
	case ErrorUserInput:
		return http.StatusBadRequest // 400
	case ErrorUserCredential:
		return http.StatusUnauthorized // 401
	case ErrorInvalidValidate:
		return http.StatusBadRequest // 400

	// Server
	case ErrorGeneralServer:
		return http.StatusInternalServerError // 500

	// Database
	case ErrorGeneralDatabase:
		return http.StatusInternalServerError // 500
	case ErrorDatabaseNotFound:
		return http.StatusNotFound

	// Default
	default:
		return http.StatusBadRequest // 400
	}
}
