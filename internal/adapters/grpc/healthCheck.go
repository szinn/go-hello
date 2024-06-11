package grpc

import (
	"context"

	pb "github.com/szinn/go-hello/proto"
)

func (server *server) IsHealthy(ctx context.Context, request *pb.IsHealthyRequest) (*pb.IsHealthyResponse, error) {
	coreServices := getCoreServices(ctx)
	logger := getLogger(ctx)
	logger.Debug("grpc IsHealthy")
	
	return &pb.IsHealthyResponse{Healthy: coreServices.HealthService.IsHealthy()}, nil
}
