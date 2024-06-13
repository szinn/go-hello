package grpc

import (
	"context"
	"github.com/szinn/go-hello/internal/core"
	"github.com/szinn/go-hello/internal/logging"

	pb "github.com/szinn/go-hello/proto"
)

func (server *server) IsHealthy(ctx context.Context, request *pb.IsHealthyRequest) (*pb.IsHealthyResponse, error) {
	coreServices := core.GetCoreServices(ctx)
	logger := logging.GetLogger(ctx)
	logger.Debug("grpc IsHealthy")

	return &pb.IsHealthyResponse{Healthy: coreServices.HealthService.IsHealthy(ctx)}, nil
}
