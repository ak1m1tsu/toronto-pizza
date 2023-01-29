package service

import (
	"context"
	"time"

	dto "github.com/romankravchuk/toronto-pizza/internal/router/handlers/models"
)

type IAuthService interface {
	GetUserByPhone(ctx context.Context, phone string) (*dto.UserDTO, error)
	CreateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error)
	ValidateToken(token string, publicKey string) (interface{}, error)
	ValidatePassword(phone string, pwd string) error
}

type IProductService interface {
	GetProductByID(ctx context.Context, id string) (*dto.ProductDTO, error)
	GetProducts(ctx context.Context) ([]*dto.ProductDTO, error)
	DeleteProduct(ctx context.Context, id string) (string, error)
	UpdateProduct(ctx context.Context, id string, product *dto.UpdateProductDTO) (*dto.ProductDTO, error)
	InsertProduct(ctx context.Context, product *dto.CreateProductDTO) (*dto.ProductDTO, error)
}
