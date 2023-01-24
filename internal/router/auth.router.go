package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/romankravchuk/toronto-pizza/internal/router/handlers"
)

func NewAuthRouter() *Router {
	router := chi.NewRouter()
	authHandler := handlers.NewAuthHandler()
	router.Post("/sign-in", authHandler.HandleSignIn)
	router.Post("/sign-up", authHandler.HandleSignUp)
	router.Post("/log-out", authHandler.HandleLogOut)
	return &Router{Mux: *router}
}
