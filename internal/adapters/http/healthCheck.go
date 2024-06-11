package http

import (
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	logger := getLogger(r)
	logger.Debug("healthCheck")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Ready"))
}
