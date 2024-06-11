package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	httpServer "github.com/szinn/go-hello/internal/adapters/http"
	"github.com/szinn/go-hello/internal/core"
	"github.com/szinn/go-hello/internal/logging"
)

func main() {
	logging.Init()

	core.Init()

	port := 8080
	server := httpServer.CreateServer(port)
	slog.Info(fmt.Sprintf("Listening on :%d", port))

	sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

	httpServer.ShutdownServer(server)

    slog.Info("Graceful shutdown complete.")
}
