package main

import (
	httpServer "github.com/szinn/go-hello/internal/adapters/http"
	"github.com/szinn/go-hello/internal/core"
	"github.com/szinn/go-hello/internal/logging"
	"log/slog"
)

func main() {
	logging.Init()

	core.Init()

	serverCtx := httpServer.CreateServer(8080)

	slog.Info("Listening on :8080")

	<-serverCtx.Done()
}
