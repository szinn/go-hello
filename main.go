package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

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
	server := httpAdapter.CreateServer(port, core)

	// Connect into the OS and wait for termination
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Shutdown the server gracefully
	server.Shutdown()

	// Shutdown the core
	core.Shutdown()

	slog.Info("Graceful shutdown complete.")
}
