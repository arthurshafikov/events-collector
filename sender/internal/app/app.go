package app

import (
	"context"
	"flag"

	"github.com/arthurshafikov/events-collector/sender/internal/config"
	"github.com/arthurshafikov/events-collector/sender/internal/logger"
	"github.com/arthurshafikov/events-collector/sender/internal/services"
	server "github.com/arthurshafikov/events-collector/sender/internal/transport/http"
	"github.com/arthurshafikov/events-collector/sender/internal/transport/http/handler"
)

var envFile = *flag.String("env", "./.env", "path to the .env file")

func Run() {
	flag.Parse()
	ctx := context.Background()

	config := config.NewConfig(envFile)
	logger := logger.NewLogger()

	services := services.NewServices(services.Deps{
		Logger: logger,
		Config: config,
	})

	h := handler.NewHandler(ctx, services)

	server.NewServer(ctx, h).Serve(config.App.Port)
}
