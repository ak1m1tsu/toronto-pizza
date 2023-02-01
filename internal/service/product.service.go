package service

import (
	"context"

	"github.com/romankravchuk/toronto-pizza/internal/repository"
	"github.com/romankravchuk/toronto-pizza/internal/repository/filter"
	dbo "github.com/romankravchuk/toronto-pizza/internal/repository/models"
	dto "github.com/romankravchuk/toronto-pizza/internal/router/handlers/models"
)

type ProductService struct {
	products repository.IProductRepository
}

func NewProductService(rep repository.IProductRepository) IProductService {
	return &ProductService{products: rep}
}

func (s *ProductService) GetProductByID(ctx context.Context, id string) (*dto.ProductDTO, error) {
	product, err := s.products.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.NewProductDTO(product.ID, product.Name, product.Description, "", product.Price), nil
}

func (s *ProductService) GetProducts(ctx context.Context, pf *dto.ProductFilter, pss []*dto.ProductSort, page int) ([]*dto.ProductDTO, error) {
	repFilter := filter.NewProductFilter(pf.Name, pf.Category, pf.PriceMin, pf.PriceMax)
	repSort := s.buildSort(pss)
	products, err := s.products.GetAll(ctx, repFilter, repSort, page)
	if err != nil {
		return nil, err
	}
	var productsDTO []*dto.ProductDTO
	for _, p := range products {
		productsDTO = append(productsDTO, dto.NewProductDTO(p.ID, p.Name, p.Description, p.Category, p.Price))
	}
	return productsDTO, nil
}

func (s *ProductService) buildSort(pss []*dto.ProductSort) *filter.ProductSort {
	var opts []filter.SortOption
	for _, ps := range pss {
		opts = append(opts, filter.NewSortOption(ps.Field, ps.Order))
	}
	return filter.NewProductSort(opts)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) (string, error) {
	err := s.products.Delete(ctx, id)
	return id, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, id string, updateProduct *dto.UpdateProductDTO) (*dto.ProductDTO, error) {
	p := dbo.NewProduct(updateProduct.Name, updateProduct.Description, updateProduct.Category, updateProduct.Price)
	p, err := s.products.Update(ctx, id, p)
	if err != nil {
		return nil, err
	}
	productDto := dto.NewProductDTO(id, p.Name, p.Description, p.Category, p.Price)
	return productDto, nil
}

func (s *ProductService) InsertProduct(ctx context.Context, createProduct *dto.CreateProductDTO) (*dto.ProductDTO, error) {
	p := dbo.NewProduct(createProduct.Name, createProduct.Description, createProduct.Category, createProduct.Price)
	p, err := s.products.Insert(ctx, p)
	if err != nil {
		return nil, err
	}
	productDto := dto.NewProductDTO(p.ID, p.Name, p.Description, p.Category, p.Price)
	return productDto, nil
}
