package service

import (
	"context"

	"github.com/romankravchuk/toronto-pizza/internal/repository"
)

type ProductService struct {
	rep repository.IProductRepository
}

func NewProductService(rep repository.IProductRepository) IProductService {
	return &ProductService{rep: rep}
}

func (s *ProductService) GetProductByID(ctx context.Context, id string) (any, error) {
	return nil, nil
}

func (s *ProductService) GetProducts(ctx context.Context) (any, error) {
	return nil, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) (any, error) {
	return nil, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id string, product any) (any, error) {
	return nil, nil
}

func (s *ProductService) InsertProduct(ctx context.Context, product any) (any, error) {
	return nil, nil
}
