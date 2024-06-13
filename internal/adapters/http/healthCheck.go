package http

import (
	"net/http"
	"github.com/szinn/go-hello/internal/core"
	"github.com/szinn/go-hello/internal/logging"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	coreServices := core.GetCoreServices(ctx)
	logger := logging.GetLogger(ctx)

	logger.Debug("http healthCheckHandler")

	if coreServices.HealthService.IsHealthy(ctx) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ready"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
