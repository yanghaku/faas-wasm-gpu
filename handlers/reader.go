package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// MakeFunctionReader handler for reading functions deployed in the cluster as deployments.
func MakeFunctionReader() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}

		functions, err := controllerInstance.FuncStateList()
		if err != nil {
			errStr := fmt.Errorf("failed to get function status list %s", err.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}

		functionBytes, err := json.Marshal(functions)
		if err != nil {
			errStr := fmt.Errorf("failed to marshal function status list %s", err.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			log.Println(errStr)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(functionBytes)
	}
}
