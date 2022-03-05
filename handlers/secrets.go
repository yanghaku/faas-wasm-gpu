package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/openfaas/faas-provider/types"
	"log"
	"net/http"
)

// MakeSecretsHandler make a secret handler
func MakeSecretsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}

		switch r.Method {
		case http.MethodGet:
			listSecrets(w)
		case http.MethodPost:
			createSecret(w, r)
		case http.MethodPut:
			replaceSecret(w, r)
		case http.MethodDelete:
			deleteSecret(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// listSecrets get all secrets
func listSecrets(w http.ResponseWriter) {
	names, err := controllerInstance.SecretsClient.List()
	if err != nil {
		errStr := fmt.Errorf("secret list error: %s", err.Error()).Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	secretsBytes, err := json.Marshal(names)
	if err != nil {
		errStr := fmt.Errorf("secrets json marshal error: %s", err.Error()).Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		log.Println(errStr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(secretsBytes)
}

// createSecret create a new secret item
func createSecret(w http.ResponseWriter, r *http.Request) {
	secret := types.Secret{}
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		errStr := fmt.Errorf("secret unmarshal error: %s", err.Error()).Error()
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	err = controllerInstance.SecretsClient.Create(secret)
	if err != nil {
		errStr := fmt.Errorf("secret create error: %s", err.Error()).Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func replaceSecret(w http.ResponseWriter, r *http.Request) {
	secret := types.Secret{}
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		errStr := fmt.Errorf("secret unmarshal error: %s", err.Error()).Error()
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	err = controllerInstance.SecretsClient.Replace(secret)
	if err != nil {
		errStr := fmt.Errorf("secret update error: %s", err.Error()).Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func deleteSecret(w http.ResponseWriter, r *http.Request) {
	secret := types.Secret{}
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		errStr := fmt.Errorf("secret unmarshal error: %s", err.Error()).Error()
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	err = controllerInstance.SecretsClient.Delete(secret.Name)
	if err != nil {
		errStr := fmt.Errorf("secret delete error: %s", err.Error()).Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
