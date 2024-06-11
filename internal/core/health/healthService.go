package health

import (
	"log/slog"
)

type HealthService struct{}

func (healthService *HealthService) IsHealthy() bool {
	slog.Debug("isHealthy called")

	return true
}
