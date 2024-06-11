package grpc

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/google/uuid"
	"github.com/szinn/go-hello/internal/core"
	pb "github.com/szinn/go-hello/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	server *grpc.Server
}

type server struct {
	pb.UnimplementedGoHelloServer
	core *core.CoreServices
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
)

type requestIdKey struct{}
type loggerKey struct{}
type coreServicesKey struct{}

func CreateServer(port string, core *core.CoreServices) *Server {
	slog.Debug("Creating GRPC server...")

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	pb.RegisterGoHelloServer(s, &server{core: core})

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatalf("GRPC server error: %v", err)
		}

		slog.Info(fmt.Sprintf("Listening for GRPC on :%s", port))
		if err := s.Serve(listener); err != nil {
			log.Fatalf("GRPC server error: %v", err)
		}
		slog.Info("Stopped serving new GRPC connections")

	}()

	slog.Debug("...GRPC server created")

	return &Server{server: s}
}

func (server *Server) Shutdown() {
	server.server.GracefulStop()
}

func interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	var id string
	ids := md["x-request-id"]
	if len(ids) == 0 {
		id = uuid.New().String()
	} else {
		id = ids[0]
	}

	ctx = context.WithValue(ctx, coreServicesKey{}, info.Server.(*server).core)
	ctx = context.WithValue(ctx, requestIdKey{}, id)
	ctx = context.WithValue(ctx, loggerKey{}, slog.With(
		slog.String("request-id", id),
	))

	return handler(ctx, req)
}

func getLogger(ctx context.Context) *slog.Logger {
	return ctx.Value(loggerKey{}).(*slog.Logger)
}

func getCoreServices(ctx context.Context) *core.CoreServices {
	return ctx.Value(coreServicesKey{}).(*core.CoreServices)
}
