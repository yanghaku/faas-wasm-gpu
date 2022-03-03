package handlers

import "net/http"

// MakeSecretsHandler make a secret handler
func MakeSecretsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
