package models

type Product struct {
	ID          string  `bson:"_id,omitempty"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float64 `bson:"price"`
	Category    string  `bson:"category"`
}

func NewProduct(name, description, category string, price float64) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
	}
}
