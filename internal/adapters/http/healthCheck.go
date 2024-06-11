package http

import (
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r)
	logger.Debug("healthCheck")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}
