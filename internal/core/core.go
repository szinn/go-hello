package core

import (
	"log/slog"
)

func Init() {
	logger := slog.Default().With(
		slog.String("tag", "value"),
	)

	logger.Debug("Core initialized")
}
