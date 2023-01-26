package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/romankravchuk/toronto-pizza/internal/config"
	"github.com/romankravchuk/toronto-pizza/internal/repository"
	"github.com/romankravchuk/toronto-pizza/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	chi.Mux
}

func NewRouter(db *mongo.Database, config *config.Config) *Router {
	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(time.Second*60),
	)

	userRep := repository.NewUserRepository(db)
	authSvc := service.NewAuthService(userRep)
	router.Mount("/auth", NewAuthRouter(authSvc, config))

	productRep := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRep)
	router.Mount("/admin", NewProductRouter(productSvc, authSvc, config))
	return &Router{Mux: *router}
}
