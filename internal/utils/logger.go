// internal/utils/logger.go
package utils

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func InitLogger() {
	Log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

func InitVerboseLogger() {
	Log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}