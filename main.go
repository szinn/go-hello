package main

import (
	"fmt"
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
	core := core.Init()

	// Initialize the HTTP server
	port := 8080
	server := httpAdapter.CreateServer(port, core)
	slog.Info(fmt.Sprintf("Listening on :%d", port))

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
