package core

import (
	"log/slog"

	healthService "github.com/szinn/go-hello/internal/core/health"
)

type CoreServices struct {
	HealthService *healthService.HealthService
}

func CreateCore() *CoreServices {
	slog.Debug("Starting core services...")

	healthService := &healthService.HealthService{}

	slog.Debug("...core services started")

	return &CoreServices{
		HealthService: healthService,
	}
}

func (core *CoreServices) Shutdown() {
	slog.Debug("Core shutdown")
}
