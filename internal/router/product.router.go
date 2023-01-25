package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/romankravchuk/toronto-pizza/internal/router/handlers"
	"github.com/romankravchuk/toronto-pizza/internal/service"
)

func NewProductRouter(svc service.IProductService) *Router {
	router := chi.NewRouter()
	productHandler := handlers.NewProductHandler(svc)
	router.Route("/product", func(r chi.Router) {
		r.With().Get("/", productHandler.HandleGetProducts)
		r.Post("/", productHandler.HandleAddProduct)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(productHandler.Context)
			r.Get("/", productHandler.HandleGetProduct)
			r.Put("/", productHandler.HandleUpdateProduct)
			r.Delete("/", productHandler.HandleDeleteProduct)
		})
	})
	return &Router{Mux: *router}
}
