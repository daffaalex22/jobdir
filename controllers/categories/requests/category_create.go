package requests

import (
	"main.go/business/categories"
)

type CategoryCreate struct {
	Category string `json:"category"`
	// Jobs     []requests.JobCreate `json:jobs`
	// CategoryDescription string `json:"categoryDesc"`
}

func (category *CategoryCreate) ToDomain() categories.Domain {
	return categories.Domain{
		Category: category.Category,
		// Jobs:     requests.ListToDomain(category.Jobs),
		// CategoryDescription: category.CategoryDescription,
	}
}
