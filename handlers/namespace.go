package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// NamespaceLister return an empty list
func NamespaceLister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}

		var namespaces []string

		out, err := json.Marshal(namespaces)
		if err != nil {
			errStr := fmt.Errorf("failed to list namespaces: %s", err.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			log.Println(errStr)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	}
}
