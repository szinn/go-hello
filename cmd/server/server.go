package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcAdapter "github.com/szinn/go-hello/internal/adapters/grpc"
	httpAdapter "github.com/szinn/go-hello/internal/adapters/http"
	"github.com/szinn/go-hello/internal/core"
	"github.com/szinn/go-hello/internal/logging"
)

func main() {
	logging.Init()

	// Initialize the core application
	core := core.CreateCore()

	// Initialize the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	httpServer := httpAdapter.CreateServer(port, core)

	// Initialize the GRPC server
	grpcServer := grpcAdapter.CreateServer("6952", core)

	// Connect into the OS and wait for termination
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Shutdown the server gracefully
	httpServer.Shutdown()
	grpcServer.Shutdown()

	// Shutdown the core
	core.Shutdown()
	time.Sleep(2 * time.Second)

	slog.Info("Graceful shutdown complete.")
}
