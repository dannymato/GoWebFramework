package error

import (
	"encoding/json"
	"net/http"
)

// Errors contains an array of Error structs for returning multiple errors to the client
type Errors struct {
	Errors []*Error `json:"errors"`
}

// Error contains the information to return to the client when an error occurs
type Error struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var (
	// ErrNotAcceptable provides a constant for HTTP Error 406
	ErrNotAcceptable = &Error{"not_acceptable", 406, "Not Acceptable", "Accept header must be set to 'application/vnd.api+json'."}
	// ErrInternalServer provides a constant Error for HTTP Error 500
	ErrInternalServer = &Error{"internal_server_error", 500, "Internal Server Error", "Something went wrong."}
	// ErrBadRequest provides a constant Error for HTTP Error 400
	ErrBadRequest = &Error{"bad_request", 400, "Bad Request Error", "Invalid JSON input"}
)

// WriteError takes an httpResponseWriter and an Error and writes the error the Writer using the JSON encoder
func WriteError(w http.ResponseWriter, err *Error) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(Errors{[]*Error{err}})
}
