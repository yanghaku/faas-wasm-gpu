package handlers

import "net/http"

// MakeReplicaUpdater updates desired count of replicas
func MakeReplicaUpdater() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// MakeReplicaReader reads the amount of replicas for a deployment
func MakeReplicaReader() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
