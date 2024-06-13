package health

import (
	"context"

	"github.com/szinn/go-hello/internal/logging"
)

type HealthService struct{}

func (healthService *HealthService) IsHealthy(ctx context.Context) bool {
	logger := logging.GetLogger(ctx)

	logger.Debug("isHealthy called")

	return true
}
