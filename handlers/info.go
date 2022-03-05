package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/openfaas/faas-provider/types"
	"log"
	"net/http"
)

const (
	//OrchestrationIdentifier identifier string for provider orchestration
	OrchestrationIdentifier = "WebAssembly"
	//ProviderName name of the provider
	ProviderName = "faas-wasm-cuda"
)

//MakeInfoHandler creates handler for /system/info endpoint
func MakeInfoHandler(version, sha string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}

		infoResponse := types.ProviderInfo{
			Orchestration: OrchestrationIdentifier,
			Name:          ProviderName,
			Version: &types.VersionInfo{
				Release: version,
				SHA:     sha,
			},
		}

		jsonOut, marshalErr := json.Marshal(infoResponse)
		if marshalErr != nil {
			errStr := fmt.Errorf("info json marshal error: %s", marshalErr.Error()).Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			log.Println(errStr)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonOut)
	}
}
