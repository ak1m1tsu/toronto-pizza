package repository

import (
	"context"

	"github.com/romankravchuk/toronto-pizza/internal/repository/models"
)

type IUserRepository interface {
	GetByPhone(ctx context.Context, phone string) (*models.User, error)
}

type IProductRepository interface {
	GetByID(ctx context.Context, id string) (*models.Product, error)
	GetAll(ctx context.Context) ([]*models.Product, error)
	Insert(ctx context.Context, user *models.Product) (*models.Product, error)
	Update(ctx context.Context, id string, user *models.Product) (*models.Product, error)
	Delete(ctx context.Context, id string) bool
}
