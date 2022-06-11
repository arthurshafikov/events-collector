package app

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/arthurshafikov/events-collector/sender/internal/config"
	"github.com/arthurshafikov/events-collector/sender/internal/logger"
	"github.com/arthurshafikov/events-collector/sender/internal/services"
	server "github.com/arthurshafikov/events-collector/sender/internal/transport/http"
	"github.com/arthurshafikov/events-collector/sender/internal/transport/http/handler"
)

func Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config := config.NewConfig()
	logger := logger.NewLogger()

	services := services.NewServices(services.Deps{
		Context: ctx,
		Logger:  logger,
		Config:  config,
	})

	h := handler.NewHandler(ctx, services)

	server.NewServer(ctx, h).Serve(config.App.Port)
}
