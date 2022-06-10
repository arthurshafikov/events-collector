package services

import (
	"context"

	"github.com/arthurshafikov/events-collector/storage/internal/config"
	"github.com/arthurshafikov/events-collector/storage/internal/core"
	"github.com/arthurshafikov/events-collector/storage/internal/repository"
)

type Collectors interface {
	StoreEvent(ctx context.Context, event core.Event) error
}

type Services struct {
	Collectors
}

type Logger interface {
	Error(err error)
	Info(msg string)
}

type Deps struct {
	Context    context.Context
	Repository *repository.Repository
	Logger     Logger
	Config     *config.Config
}

func NewServices(deps Deps) *Services {
	return &Services{
		Collectors: NewCollector(deps.Logger, deps.Repository.Collector, deps.Config.App.BufferSizeLimit),
	}
}
