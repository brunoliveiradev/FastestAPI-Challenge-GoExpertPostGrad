package util

import (
	"errors"
	"net"
	"net/http"
)

type CustomError struct {
	Status  int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func HandleError(w http.ResponseWriter, err error) {
	var nerr net.Error
	var cerr *CustomError

	switch {
	case errors.As(err, &nerr) && nerr.Timeout():
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
	case errors.As(err, &cerr):
		http.Error(w, cerr.Message, cerr.Status)
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
