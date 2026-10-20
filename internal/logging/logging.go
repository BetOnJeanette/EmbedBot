// Package logging contains helper functions for logging
package logging

import (
	"log/slog"
	"os"
)

// Init initializes the logger.
func Init() {
	var level = new(slog.LevelVar)
	level.Set(slog.LevelDebug)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	slog.SetDefault(slog.New(h))
}
