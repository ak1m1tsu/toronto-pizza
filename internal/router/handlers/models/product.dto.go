package models

import "errors"

var (
	ErrInvalidProductName     = errors.New("invalid product name")
	ErrInvalidProductCategory = errors.New("invalid product category")
	ErrInvalidProductPrice    = errors.New("invalid product price")
)

type ProductDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

type CreateProductDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

type UpdateProductDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

func ValidateCreateProductDTO(dto *CreateProductDTO) error {
	if len(dto.Name) < 3 {
		return ErrInvalidProductName
	}
	if dto.Category == "" || len(dto.Category) == 0 {
		return ErrInvalidProductCategory
	}
	if dto.Price < 0 {
		return ErrInvalidProductPrice
	}
	return nil
}

func NewProductDTO(id, name, description, category string, price float64) *ProductDTO {
	return &ProductDTO{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
	}
}
