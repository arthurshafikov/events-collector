package handler

import (
	"context"

	"github.com/arthurshafikov/events-collector/sender/internal/services"
	"github.com/fasthttp/router"
)

type Handler struct {
	ctx      context.Context
	services *services.Services
}

func NewHandler(
	ctx context.Context,
	services *services.Services,
) *Handler {
	return &Handler{
		ctx:      ctx,
		services: services,
	}
}

func (h *Handler) Init(r *router.Router) {
	h.initAuthRoutes(r)
}
