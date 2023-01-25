package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/romankravchuk/toronto-pizza/internal/repository"
	"github.com/romankravchuk/toronto-pizza/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	chi.Mux
}

func NewRouter(db *mongo.Database) *Router {
	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(time.Second*60),
	)

	userRep := repository.NewUserRepository(db)
	authSvc := service.NewAuthService(userRep)
	router.Mount("/auth", NewAuthRouter(authSvc))

	productRep := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRep)
	router.Mount("/admin", NewProductRouter(productSvc))
	return &Router{Mux: *router}
}
