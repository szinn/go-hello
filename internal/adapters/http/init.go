package http

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

const requestIdHeader = "x-request-id"

type requestIdKey struct{}
type loggerKey struct{}

func CreateServer(port int) context.Context {
	mux := http.NewServeMux()
	mux.HandleFunc("/-/ready", healthCheck)

	ctx, cancelCtx := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: HandlerWrapper(mux),
	}
	go func() {
		server.ListenAndServe()
		cancelCtx()
	}()

	return ctx
}

func HandlerWrapper(h http.Handler) http.Handler {
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

		h.ServeHTTP(w, req.WithContext(ctx))
	})
}

func GetRequestId(r *http.Request) string {
	return r.Context().Value(requestIdKey{}).(string)
}

func GetLogger(r *http.Request) *slog.Logger {
	return r.Context().Value(loggerKey{}).(*slog.Logger)
}
