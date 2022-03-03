package handlers

import "net/http"

// MakeDeployHandler creates a handler to create new functions in the cluster
func MakeDeployHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
