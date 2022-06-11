package handler

import (
	"context"
	"encoding/json"

	"github.com/arthurshafikov/events-collector/sender/internal/services"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
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
}

func (h *Handler) setJSONResponse(ctx *fasthttp.RequestCtx, body interface{}) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		ctx.Error(err.Error(), 500)
		return
	}
	ctx.SetBody(bodyJSON)
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
}
