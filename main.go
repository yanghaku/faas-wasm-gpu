package main

import (
	faasProvider "github.com/openfaas/faas-provider"
	"github.com/openfaas/faas-provider/logs"
	providerTypes "github.com/openfaas/faas-provider/types"
	"github.com/yanghaku/faas-wasm-gpu/handlers"
	"github.com/yanghaku/faas-wasm-gpu/version"
	"log"
	"os"
	"strconv"
	"time"
)

func init() {
	f, ok := os.LookupEnv("log_file")
	if ok { // if not set log_file, use the stdout
		file, err := os.Create(f)
		if err != nil {
			log.SetOutput(file)
		}
	}

	log.SetPrefix("[faas-wasm-cuda] ")
	log.SetFlags(log.Ltime | log.Lmicroseconds)
}

func main() {

	sha, release := version.GetReleaseInfo()
	log.Printf("faas-wasm-cuda version:%s. Last commit SHA: %s\n", release, sha)

	defaultTCPPort := 8081
	defaultReadTimeout := 3
	defaultWriteTimeout := 3

	readTimeout := time.Duration(parseIntValue(os.Getenv("read_timeout"), defaultReadTimeout)) * time.Second
	writeTimeout := time.Duration(parseIntValue(os.Getenv("write_timeout"), defaultWriteTimeout)) * time.Second
	port := parseIntValue(os.Getenv("port"), defaultTCPPort)
	log.Printf("tcp port = %d\n", port)

	log.Println("Starting controller")
	runController(port, readTimeout, writeTimeout)

}

func runController(port int, readTimeout time.Duration, writeTimeout time.Duration) {

	bootstrapHandlers := providerTypes.FaaSHandlers{
		FunctionProxy:        handlers.MakeProxy(),
		DeleteHandler:        handlers.MakeDeleteHandler(),
		DeployHandler:        handlers.MakeDeployHandler(),
		FunctionReader:       handlers.MakeFunctionReader(),
		ReplicaReader:        handlers.MakeReplicaReader(),
		ReplicaUpdater:       handlers.MakeReplicaUpdater(),
		UpdateHandler:        handlers.MakeUpdateHandler(),
		HealthHandler:        handlers.MakeHealthHandler(),
		InfoHandler:          handlers.MakeInfoHandler(version.BuildVersion(), version.GitCommit),
		SecretHandler:        handlers.MakeSecretsHandler(),
		LogHandler:           logs.NewLogHandlerFunc(handlers.NewLogRequester(), writeTimeout),
		ListNamespaceHandler: handlers.NamespaceLister(),
	}

	bootstrapConfig := providerTypes.FaaSConfig{
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
		TCPPort:         &port,
		EnableBasicAuth: false,
	}

	faasProvider.Serve(&bootstrapHandlers, &bootstrapConfig)
}

func parseIntValue(val string, fallback int) int {
	if len(val) > 0 {
		parsedVal, parseErr := strconv.Atoi(val)
		if parseErr == nil && parsedVal >= 0 {
			return parsedVal
		}
	}
	return fallback
}
