package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/szinn/go-hello/internal/core"
	"github.com/szinn/go-hello/internal/logging"
)

type Server struct {
	server *http.Server
}

const requestIdHeader = "x-request-id"

type requestIdKey struct{}

func CreateServer(port string, core *core.CoreServices) *Server {
	slog.Debug("Creating HTTP server...")
	mux := http.NewServeMux()
	addHandlers(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handlerWrapper(mux, core),
	}

	go func() {
		slog.Info(fmt.Sprintf("Listening for HTTP on :%s", port))
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		slog.Info("Stopped serving new HTTP connections")
	}()

	slog.Debug("...HTTP server created")

	return &Server{server}
}

func (server *Server) Shutdown() {
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	_ = server.server.Shutdown(shutdownCtx)
}

func handlerWrapper(h http.Handler, coreServices *core.CoreServices) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Ensure the request has an id
		id := req.Header.Get(requestIdHeader)
		if id == "" {
			id = uuid.New().String()
			req.Header.Set(requestIdHeader, id)
		}

		// Add core services, id, and logger to context
		ctx := req.Context()
		ctx = context.WithValue(ctx, requestIdKey{}, id)
		ctx = core.SetCoreServices(ctx, coreServices)
		ctx = logging.SetLogger(ctx, slog.With(
			slog.String("request-id", id),
		))

		// Add id to response
		w.Header().Set(requestIdHeader, id)

		// Serve the request
		h.ServeHTTP(w, req.WithContext(ctx))
	})
}

// func getRequestId(r *http.Request) string {
// 	return r.Context().Value(requestIdKey{}).(string)
// }
