package services

import (
	"context"

	"github.com/arthurshafikov/events-collector/sender/internal/client/generated"
	"github.com/arthurshafikov/events-collector/sender/internal/core"
)

type AuthService struct {
	logger Logger
	client generated.CollectorClient
}

func NewAuthService(logger Logger, client generated.CollectorClient) *AuthService {
	return &AuthService{
		logger: logger,
		client: client,
	}
}

func (a *AuthService) Register(ctx context.Context, user core.User) error {
	// call to repository, store user
	a.client.StoreEvent(ctx, &generated.EventRequest{
		EventType: "registration",
	})

	return nil
}

func (a *AuthService) Login(ctx context.Context, inp core.AuthorizeInput) error {
	// call to repository, check if user exists
	a.client.StoreEvent(ctx, &generated.EventRequest{
		EventType: "login",
	})

	return nil
}
