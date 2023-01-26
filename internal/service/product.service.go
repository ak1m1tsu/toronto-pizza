package service

import (
	"context"

	"github.com/romankravchuk/toronto-pizza/internal/repository"
	dto "github.com/romankravchuk/toronto-pizza/internal/router/handlers/models"
)

type ProductService struct {
	rep repository.IProductRepository
}

func NewProductService(rep repository.IProductRepository) IProductService {
	return &ProductService{rep: rep}
}

func (s *ProductService) GetProductByID(ctx context.Context, id string) (*dto.ProductDTO, error) {
	return nil, nil
}
func (s *ProductService) GetProducts(ctx context.Context) ([]*dto.ProductDTO, error) {
	return nil, nil
}
func (s *ProductService) DeleteProduct(ctx context.Context, id string) (string, error) {
	return "", nil
}
func (s *ProductService) UpdateProduct(ctx context.Context, id string, product *dto.ProductDTO) (*dto.ProductDTO, error) {
	return nil, nil
}
func (s *ProductService) InsertProduct(ctx context.Context, product *dto.ProductDTO) (*dto.ProductDTO, error) {
	return nil, nil
}
