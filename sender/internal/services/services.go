package services

import (
	"context"

	"github.com/arthurshafikov/events-collector/sender/internal/client/generated"
	"github.com/arthurshafikov/events-collector/sender/internal/config"
	"github.com/arthurshafikov/events-collector/sender/internal/core"
)

type Auth interface {
	Register(ctx context.Context, user core.User) error
	Login(ctx context.Context, inp core.AuthorizeInput) error
}

type Services struct {
	Auth
}

type Logger interface {
	Error(err error)
	Info(msg string)
}

type Deps struct {
	Logger Logger
	Config *config.Config
	Client generated.CollectorClient
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth: NewAuthService(deps.Logger, deps.Client),
	}
}
