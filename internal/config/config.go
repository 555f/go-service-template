package config

import (
	"log/slog"
)

// Config
// @gg:"config"
// @cfg-md-doc
// @cfg-env-file:"./dev.env"
type Config struct {
	// Адрес сервера
	// @cfg-required
	Addr string
	// Порт сервера
	// @cfg-required
	Port int
	// Уровень логирования
	// debug, info, warn, error
	LogLevel string
}

func (cfg *Config) SLogLevel() slog.Level {
	switch cfg.LogLevel {
	default:
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}
}
