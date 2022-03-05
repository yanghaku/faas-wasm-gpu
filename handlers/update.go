package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/openfaas/faas-provider/types"
	"io/ioutil"
	"net/http"
)

// MakeUpdateHandler update specified function
func MakeUpdateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}

		body, _ := ioutil.ReadAll(r.Body)

		request := types.FunctionDeployment{}
		err := json.Unmarshal(body, &request)
		if err != nil {
			errStr := fmt.Errorf("unable to unmarshal request: %s", err.Error()).Error()
			http.Error(w, errStr, http.StatusBadRequest)
			return
		}

		err = controllerInstance.UpdateFunc(&request)
		if err != nil {
			errStr := fmt.Errorf("update function %s fail: %s", request.Service, err.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
