package models

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	ErrInvalidProductName     = errors.New("invalid product name")
	ErrInvalidProductCategory = errors.New("invalid product category")
	ErrInvalidProductPrice    = errors.New("invalid product price")
)

type ProductFilter struct {
	Name     string
	Category string
	PriceMin float64
	PriceMax float64
}

func NewProductFilter(r *http.Request) *ProductFilter {
	var min, max float64
	var err error
	min, err = strconv.ParseFloat(chi.URLParam(r, "priceMin"), strconv.IntSize)
	if err != nil {
		min = 0
	}
	max, err = strconv.ParseFloat(chi.URLParam(r, "priceMax"), strconv.IntSize)
	if err != nil {
		max = min
	}
	return &ProductFilter{
		Name:     chi.URLParam(r, "name"),
		Category: chi.URLParam(r, "category"),
		PriceMin: min,
		PriceMax: max,
	}
}

type ProductSort struct {
	Field string
	Order int
}

func NewProductSort(r *http.Request) *ProductSort {
	return &ProductSort{}
}

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
