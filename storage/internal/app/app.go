package app

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/arthurshafikov/events-collector/storage/internal/config"
	"github.com/arthurshafikov/events-collector/storage/internal/logger"
	"github.com/arthurshafikov/events-collector/storage/internal/repository"
	"github.com/arthurshafikov/events-collector/storage/internal/services"
	grpcapi "github.com/arthurshafikov/events-collector/storage/internal/transport/grpc/api"
	"github.com/arthurshafikov/events-collector/storage/pkg/clickhousedb"
)

func Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config := config.NewConfig()
	logger := logger.NewLogger()

	conn, err := clickhousedb.NewClickhouseDB(ctx, clickhousedb.Options{
		Address:  config.DB.Address,
		Database: config.DB.Database,
		Username: config.DB.Username,
		Password: config.DB.Password,
	})
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(conn)
	services := services.NewServices(services.Deps{
		Context:    ctx,
		Repository: repo,
		Logger:     logger,
		Config:     config,
	})

	grpcapi.RunGrpcServer(ctx, config.App.Port, logger, services)
}
