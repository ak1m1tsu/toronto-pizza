package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/romankravchuk/toronto-pizza/internal/config"
	"github.com/romankravchuk/toronto-pizza/internal/router/handlers"
	"github.com/romankravchuk/toronto-pizza/internal/service"
)

func NewAuthRouter(svc service.IAuthService, config *config.Config) *Router {
	router := chi.NewRouter()
	authHandler := handlers.NewAuthHandler(svc, config.AccessToken, config.RefreshToken)
	router.Post("/sign-in", authHandler.HandleSignIn)
	router.Post("/sign-up", authHandler.HandleSignUp)
	router.Post("/log-out", authHandler.HandleLogOut)
	router.Post("/refresh-token", authHandler.HandleRefreshToken)
	return &Router{Mux: *router}
}
