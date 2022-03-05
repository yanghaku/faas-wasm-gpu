package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/openfaas/faas-provider/types"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// MakeReplicaUpdater updates desired count of replicas
func MakeReplicaUpdater() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		functionName := vars["name"]

		req := types.ScaleServiceRequest{}

		if r.Body != nil {
			defer r.Body.Close()
			bytesIn, _ := ioutil.ReadAll(r.Body)
			marshalErr := json.Unmarshal(bytesIn, &req)
			if marshalErr != nil {
				errStr := fmt.Errorf("cannot parse request: %s", marshalErr.Error()).Error()
				http.Error(w, errStr, http.StatusBadRequest)
				log.Println(errStr)
				return
			}
		}

		replicas := int32(req.Replicas)
		err := controllerInstance.ScaleFunc(functionName, replicas)
		if err != nil {
			errStr := fmt.Errorf("unable to update function deployment %s: %s", functionName, err.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

// MakeReplicaReader reads the amount of replicas for a deployment
func MakeReplicaReader() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}

		vars := mux.Vars(r)
		functionName := vars["name"]

		s := time.Now()

		function, err := controllerInstance.FuncState(functionName)
		if err != nil {
			errStr := fmt.Errorf("unable to fetch function %s: %s", functionName, err.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			log.Println(errStr)
			return
		}

		if function == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		d := time.Since(s)
		log.Printf("Replicas: %s, (%d/%d) %dms\n", functionName, function.AvailableReplicas, function.Replicas, d.Milliseconds())

		functionBytes, err := json.Marshal(function)
		if err != nil {
			errStr := fmt.Errorf("failed to marshal function: %s", err.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			log.Println(errStr)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(functionBytes)
	}
}
