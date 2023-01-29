package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/romankravchuk/toronto-pizza/internal/config"
	"github.com/romankravchuk/toronto-pizza/internal/router/handlers"
	"github.com/romankravchuk/toronto-pizza/internal/service"
)

func NewAuthRouter(svc service.IAuthService, config *config.Config) *Router {
	r := &Router{chi.NewRouter()}
	authHandler := handlers.NewAuthHandler(svc, config.AccessToken, config.RefreshToken)
	r.Post("/sign-in", authHandler.HandleSignIn)
	r.Post("/log-out", authHandler.HandleLogOut)
	r.Post("/refresh-token", authHandler.HandleRefreshToken)
	return r
}
