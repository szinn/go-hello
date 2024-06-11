package main

import (
	httpServer "github.com/szinn/go-hello/internal/adapters/http"
	"github.com/szinn/go-hello/internal/core"
	"github.com/szinn/go-hello/internal/logging"
	"log/slog"
	"rsc.io/quote"
)

func main() {
	logging.Init()

	slog.Info(quote.Opt())

	core.Init()

	serverCtx := httpServer.CreateServer(8080)
	<-serverCtx.Done()
}
