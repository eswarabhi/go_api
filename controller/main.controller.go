package controller

import (
	"net/http"
)

// ConfigureRoutes will configure routes for the whole app
func ConfigureRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthcheck", HealthCheck)
	mux.HandleFunc("/providers", GetProviders)
	mux.HandleFunc("/connectors", GetConnectors)
}
