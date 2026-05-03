package server

import (
	"log/slog"
	"net/http"
)

type ErrorResponse func(w http.ResponseWriter, r *http.Request) error

func HandleErrors(e ErrorResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := e(w, r); err != nil {
			slog.Error("Internal Server Error", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
