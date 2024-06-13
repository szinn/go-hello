package logging

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

type loggerKey struct{}

func Init() {
	level := getLevel()
	handler := getHandler(level)
	logger := slog.New(handler)

	slog.SetDefault(logger)
}

func getLevel() slog.Level {
	var level slog.Level

	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	return level
}

func getHandler(level slog.Level) slog.Handler {
	handlerOptions := &slog.HandlerOptions{
		// AddSource: true,
		Level: level,
	}

	var handler slog.Handler
	if os.Getenv("LOG_FORMAT") == "text" {
		handler = slog.NewTextHandler(os.Stdout, handlerOptions)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, handlerOptions)
	}

	return handler
}

func SetLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func GetLogger(ctx context.Context) *slog.Logger {
	return ctx.Value(loggerKey{}).(*slog.Logger)
}
