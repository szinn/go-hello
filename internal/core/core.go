package core

import (
	"context"
	"log/slog"

	healthService "github.com/szinn/go-hello/internal/core/health"
)

type CoreServices struct {
	HealthService *healthService.HealthService
}

type coreServicesKey struct{}

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

func SetCoreServices(ctx context.Context, core *CoreServices) context.Context {
	return context.WithValue(ctx, coreServicesKey{}, core)
}

func GetCoreServices(ctx context.Context) *CoreServices {
	return ctx.Value(coreServicesKey{}).(*CoreServices)
}
