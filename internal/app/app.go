package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/555f/go-service-template/internal/api/controller"
	"github.com/555f/go-service-template/internal/api/presenter"
	"github.com/555f/go-service-template/internal/api/server"

	"github.com/555f/go-service-template/internal/config"

	"github.com/555f/go-service-template/internal/infrastructure/gateway"

	"github.com/555f/go-service-template/internal/usecase/user"
)

func Run(cfg *config.Config, logger *slog.Logger) error {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	registry := prometheus.NewRegistry()
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	userGateway := gateway.NewUserGateway()
	userPresenter := presenter.NewUserPresenter()
	searchUserUseCase := user.NewDefaultSearchUserUseCase(userGateway, userPresenter)

	userController := controller.NewUserController(searchUserUseCase)

	r := chi.NewRouter()

	server.SetupRoutesUserController(userController, r)

	r.Get("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}).ServeHTTP)

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           r,
	}

	logger.Debug("service started", "addr", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error("start server", "err", err)
		}
	}()

	<-sigint

	return nil
}
