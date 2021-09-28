package responses

import (
	"time"

	"main.go/business/categories"
)

type CategoryResponse struct {
	Id        int       `json:"id"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain categories.Domain) CategoryResponse {
	return CategoryResponse{
		Id:        domain.Id,
		Category:  domain.Category,
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
