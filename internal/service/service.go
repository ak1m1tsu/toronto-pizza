package service

import (
	"context"
	"time"
)

type IAuthService interface {
	GetUserByPhone(ctx context.Context, phone string) (any, error)
	CreateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error)
	ValidateToken(token string, publicKey string) (interface{}, error)
	ValidatePassword(id string, pwd string) error
}

type IProductService interface {
	GetProductByID(ctx context.Context, id string) (any, error)
	GetProducts(ctx context.Context) (any, error)
	DeleteProduct(ctx context.Context, id string) (any, error)
	UpdateProduct(ctx context.Context, id string, product any) (any, error)
	InsertProduct(ctx context.Context, product any) (any, error)
}
