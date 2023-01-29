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
	*chi.Mux
}

func NewRouter(db *mongo.Database, config *config.Config) *Router {
	r := &Router{chi.NewRouter()}
	r.SetupMiddlewares()
	r.MountAuthRouter(db, config)
	r.MountProductRouter(db, config)
	return r
}

func (r *Router) SetupMiddlewares() {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))
}

func (r *Router) MountAuthRouter(db *mongo.Database, config *config.Config) {
	rep := repository.NewUserRepository(db)
	svc := service.NewAuthService(rep)
	r.Mount("/auth", NewAuthRouter(svc, config))
}

func (r *Router) MountProductRouter(db *mongo.Database, config *config.Config) {
	uRep := repository.NewUserRepository(db)
	aSvc := service.NewAuthService(uRep)
	pRep := repository.NewProductRepository(db)
	pSvc := service.NewProductService(pRep)
	r.Mount("/admin", NewProductRouter(pSvc, aSvc, config))
}
