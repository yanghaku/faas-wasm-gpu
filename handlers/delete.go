package handlers

import "net/http"

// MakeDeleteHandler delete a function
func MakeDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
