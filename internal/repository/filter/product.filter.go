package filter

type ProductSort struct {
	Options []SortOption
}

func NewProductSort(options []SortOption) *ProductSort {
	return &ProductSort{
		Options: options,
	}
}

type ProductFilter struct {
	Name     string
	Category string
	PriceMin float64
	PriceMax float64
}

func NewProductFilter(name, category string, priceMin, priceMax float64) *ProductFilter {
	return &ProductFilter{
		Name:     name,
		Category: category,
		PriceMin: priceMin,
		PriceMax: priceMax,
	}
}
