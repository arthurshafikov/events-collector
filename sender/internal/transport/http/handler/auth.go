package handler

import (
	"encoding/json"

	"github.com/arthurshafikov/events-collector/sender/internal/core"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func (h *Handler) initAuthRoutes(r *router.Router) {
	r.POST("/register", h.handleRegister)
	r.POST("/login", h.handleLogin)
}

func (h *Handler) handleRegister(ctx *fasthttp.RequestCtx) {
	var user core.User

	if err := json.Unmarshal(ctx.Request.Body(), &user); err != nil {
		ctx.SetStatusCode(500)
		return
	}

	if err := h.services.Auth.Register(h.ctx, user); err != nil {
		ctx.SetStatusCode(500)
		return
	}

	ctx.SetStatusCode(200)
}

func (h *Handler) handleLogin(ctx *fasthttp.RequestCtx) {
	var inp core.AuthorizeInput

	if err := json.Unmarshal(ctx.Request.Body(), &inp); err != nil {
		ctx.SetStatusCode(500)
		return
	}

	if err := h.services.Auth.Login(h.ctx, inp); err != nil {
		ctx.SetStatusCode(500)
		return
	}

	ctx.SetStatusCode(200)
}
