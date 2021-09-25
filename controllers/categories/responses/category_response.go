package responses

import (
	"time"

	"main.go/business/categories"
	"main.go/business/jobs"
	"main.go/controllers/jobs/responses"
)

type CategoryResponse struct {
	Id       int                     `json:"id"`
	Category string                  `json:"category"`
	Jobs     []responses.JobResponse `json:"jobs"`
	// CategoryDescription string    `json:"categoryDesc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain categories.Domain) CategoryResponse {
	return CategoryResponse{
		Id:       domain.Id,
		Category: domain.Category,
		Jobs:     responses.ListFromDomain(domain.Jobs),
		// CategoryDescription: domain.CategoryDescription,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ListFromDomain(domain []categories.Domain) (response []CategoryResponse) {
	for _, Category := range domain {
		response = append(response, FromDomain(Category))
	}
	return
}

func ToJobsCategoryDomain(input categories.Domain) jobs.CategoryDomain {
	var category jobs.CategoryDomain
	category.Id = input.Id
	category.Category = input.Category
	category.Jobs = input.Jobs
	category.CreatedAt = input.CreatedAt
	category.UpdatedAt = input.UpdatedAt

	return category
}

func ToListJobsCategoryDomain(categories []categories.Domain) (domain []jobs.CategoryDomain) {
	for _, category := range categories {
		domain = append(domain, ToJobsCategoryDomain(category))
	}
	return
}
