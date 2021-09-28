package requests

import (
	"main.go/business/categories"
)

type CategoryCreate struct {
	Category string `json:"category"`
}

func (category *CategoryCreate) ToDomain() categories.Domain {
	return categories.Domain{
		Category: category.Category,
	}
}
