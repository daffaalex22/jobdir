package responses

import (
	"time"

	"main.go/business/applications"
)

type ApplicationResponse struct {
	UserId    int       `json:"UserId"`
	JobId     int       `json:"jobId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain applications.Domain) ApplicationResponse {
	return ApplicationResponse{
		UserId:    domain.UserId,
		JobId:     domain.JobId,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ListFromDomain(domain []applications.Domain) (response []ApplicationResponse) {
	for _, Application := range domain {
		response = append(response, FromDomain(Application))
	}
	return
}

// func ToCategoryDomain(input Applications.CategoryDomain) categories.Domain {
// 	var category categories.Domain
// 	category.Id = input.Id
// 	category.Category = input.Category
// 	category.Applications = input.Applications
// 	category.CreatedAt = input.CreatedAt
// 	category.UpdatedAt = input.UpdatedAt

// 	return category
// }

// func ToListCategoryDomain(categories []Applications.CategoryDomain) (domain []categories.Domain) {
// 	for _, category := range categories {
// 		domain = append(domain, ToCategoryDomain(category))
// 	}
// 	return
// }
