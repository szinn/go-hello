package http

import (
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	coreServices := getCoreServices(r)
	logger := getLogger(r)
	logger.Debug("healthCheck")

	if coreServices.HealthService.IsHealthy() {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ready"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
