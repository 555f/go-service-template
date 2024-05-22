package main

import (
	"log/slog"
	"os"

	"github.com/555f/go-service-template/internal/app"

	"github.com/555f/go-service-template/internal/config"
)

func main() {
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelError)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}))
	slog.SetDefault(logger)

	cfg, err := config.New()
	if err != nil {
		logger.Error("load config", "err", err)
		return
	}

	lvl.Set(cfg.SLogLevel())

	if err := app.Run(cfg, logger); err != nil {
		logger.Error("run service", "err", err)
		return
	}
}
