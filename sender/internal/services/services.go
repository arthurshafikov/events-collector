package services

import (
	"context"

	"github.com/arthurshafikov/events-collector/sender/internal/config"
)

type Services struct {
}

type Logger interface {
	Error(err error)
	Info(msg string)
}

type Deps struct {
	Context context.Context
	Logger  Logger
	Config  *config.Config
}

func NewServices(deps Deps) *Services {
	return &Services{}
}
