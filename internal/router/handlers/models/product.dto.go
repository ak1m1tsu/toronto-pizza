package models

type ProductDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
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
