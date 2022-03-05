package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/openfaas/faas-provider/types"
	"github.com/yanghaku/faas-wasm-cuda/controller"
	"io/ioutil"
	"net/http"
)

// controllerInstance the global variable for handle package to visit controller package
var controllerInstance = controller.NewController()

// MakeDeployHandler creates a handler to create new functions in the cluster
func MakeDeployHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}

		body, _ := ioutil.ReadAll(r.Body)

		request := types.FunctionDeployment{}
		err := json.Unmarshal(body, &request)
		if err != nil {
			errStr := fmt.Errorf("failed to unmarshal request: %s", err.Error()).Error()
			http.Error(w, errStr, http.StatusBadRequest)
			return
		}

		err = controllerInstance.DeployFunc(&request)
		if err != nil {
			errStr := fmt.Errorf("deploy function %s fail: %s", request.Service, err.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
