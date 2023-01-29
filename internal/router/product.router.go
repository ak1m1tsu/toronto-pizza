package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/romankravchuk/toronto-pizza/internal/config"
	"github.com/romankravchuk/toronto-pizza/internal/router/handlers"
	"github.com/romankravchuk/toronto-pizza/internal/service"
)

func NewProductRouter(prodSvc service.IProductService, authSvc service.IAuthService, config *config.Config) *Router {
	r := &Router{chi.NewRouter()}
	productHandler := handlers.NewProductHandler(prodSvc)
	authMw := handlers.NewJWTAuthMiddleware(authSvc, config.AccessToken)
	r.With(authMw.JWTRequired).Route("/product", func(r chi.Router) {
		r.With().Get("/", productHandler.HandleGetProducts)
		r.Post("/", productHandler.HandleAddProduct)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(productHandler.Context)
			r.Get("/", productHandler.HandleGetProduct)
			r.Put("/", productHandler.HandleUpdateProduct)
			r.Delete("/", productHandler.HandleDeleteProduct)
		})
	})
	return r
}
