package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// NamespaceLister return an empty list
func NamespaceLister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		namespaces := []string{}
		nsJSON, err := json.Marshal(namespaces)

		if err != nil {
			log.Printf("Unable to marshal namespaces into JSON %q", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("\"error\": \"unable to return namespaces\""))
		}

		w.WriteHeader(http.StatusOK)
		w.Write(nsJSON)
	}
}
