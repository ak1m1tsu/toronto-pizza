package models

import "fmt"

type Product struct {
	ID          string  `bson:"_id,omitempty"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float64 `bson:"price"`
	CategoryID  string  `bson:"category_id"`
}

func NewProduct(name, description string, price float64, category Category) (*Product, error) {
	if category.ID == "" {
		return nil, fmt.Errorf("failed to create product: invalid category \"%s\"", category.ID)
	}
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  category.ID,
	}, nil
}
