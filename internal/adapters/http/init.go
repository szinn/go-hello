package http

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Server struct {
	server *http.Server
}

const requestIdHeader = "x-request-id"

type requestIdKey struct{}
type loggerKey struct{}

func CreateServer(port int) *Server {
	mux := http.NewServeMux()
	addHandlers(mux)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: handlerWrapper(mux),
	}
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		slog.Info("Stopped serving new connections")
	}()

	return &Server{ server} 
}

func (server *Server) ShutdownServer() {
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	_ = server.server.Shutdown(shutdownCtx)
}

func handlerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		id := req.Header.Get(requestIdHeader)
		if id == "" {
			id = uuid.New().String()
			req.Header.Set(requestIdHeader, id)
		}

		ctx := context.WithValue(req.Context(), requestIdKey{}, id)
		ctx = context.WithValue(ctx, loggerKey{}, slog.With(
			slog.String("request-id", id),
		))

		w.Header().Set(requestIdHeader, id)
		h.ServeHTTP(w, req.WithContext(ctx))
	})
}

// func getRequestId(r *http.Request) string {
// 	return r.Context().Value(requestIdKey{}).(string)
// }

func getLogger(r *http.Request) *slog.Logger {
	return r.Context().Value(loggerKey{}).(*slog.Logger)
}
