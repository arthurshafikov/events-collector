package services

import (
	"context"

	"github.com/arthurshafikov/events-collector/internal/core"
	"github.com/arthurshafikov/events-collector/internal/repository"
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
}

func NewServices(deps Deps) *Services {
	return &Services{
		Collectors: NewCollector(deps.Logger, deps.Repository.Collector, 5),
	}
}