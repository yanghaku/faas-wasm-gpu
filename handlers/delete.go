package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/openfaas/faas-provider/types"
	"io/ioutil"
	"net/http"
)

// MakeDeleteHandler delete a function
func MakeDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}

		body, _ := ioutil.ReadAll(r.Body)

		request := types.DeleteFunctionRequest{}
		err := json.Unmarshal(body, &request)
		if err != nil {
			errStr := fmt.Errorf("failed to unmarshal request: %s", err.Error()).Error()
			http.Error(w, errStr, http.StatusBadRequest)
			return
		}

		if len(request.FunctionName) == 0 {
			http.Error(w, "function name should not be empty", http.StatusBadRequest)
			return
		}

		err = controllerInstance.DeleteFunc(request.FunctionName)
		if err != nil {
			errStr := fmt.Errorf("failed to delete function %s : %s", request.FunctionName, err.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		return
	}
}
