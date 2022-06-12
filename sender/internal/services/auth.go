package services

import (
	"context"

	"github.com/arthurshafikov/events-collector/sender/internal/core"
)

type AuthService struct {
	logger Logger
}

func NewAuthService(logger Logger) *AuthService {
	return &AuthService{
		logger: logger,
	}
}

func (a *AuthService) Register(ctx context.Context, user core.User) error {
	// call to repository, store user

	return nil
}

func (a *AuthService) Login(ctx context.Context, inp core.AuthorizeInput) error {
	// call to repository, check if user exists

	return nil
}
