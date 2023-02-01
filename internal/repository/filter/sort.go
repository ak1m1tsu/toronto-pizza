package filter

type SortOrder int

const (
	Ascending SortOrder = iota
	Descending
)

type SortOption struct {
	Field string
	Order SortOrder
}

func NewSortOption(field string, order int) SortOption {
	var sortOrder SortOrder
	switch order {
	case int(Descending):
		sortOrder = Ascending
	default:
		sortOrder = Ascending
	}
	return SortOption{
		Field: field,
		Order: sortOrder,
	}
}
