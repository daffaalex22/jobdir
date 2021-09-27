package responses

import (
	"time"

	"main.go/business/categories"
	"main.go/business/jobs"
	"main.go/controllers/applications/responses"
)

type JobResponse struct {
	Id           int                             `json:"id"`
	Title        string                          `json:"title"`
	CategoryId   int                             `json:"categoryId"`
	JobDesc      string                          `json:"jobDesc"`
	CreatedBy    int                             `json:"createdBy"`
	CompanyId    int                             `json:"companyId"`
	Applications []responses.ApplicationResponse `json:"applications"`
	CreatedAt    time.Time                       `json:"createdAt"`
	UpdatedAt    time.Time                       `json:"updatedAt"`
}

func FromDomain(domain jobs.Domain) JobResponse {
	return JobResponse{
		Id:           domain.Id,
		Title:        domain.Title,
		CategoryId:   domain.CategoryId,
		JobDesc:      domain.JobDesc,
		CreatedBy:    domain.CreatedBy,
		CompanyId:    domain.CompanyId,
		Applications: responses.ListFromDomain(domain.Applications),
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func ListFromDomain(domain []jobs.Domain) (response []JobResponse) {
	for _, job := range domain {
		response = append(response, FromDomain(job))
	}
	return
}

func ToCategoryDomain(input jobs.CategoryDomain) categories.Domain {
	var category categories.Domain
	category.Id = input.Id
	category.Category = input.Category
	category.Jobs = input.Jobs
	category.CreatedAt = input.CreatedAt
	category.UpdatedAt = input.UpdatedAt

	return category
}

func ToListCategoryDomain(categories []jobs.CategoryDomain) (domain []categories.Domain) {
	for _, category := range categories {
		domain = append(domain, ToCategoryDomain(category))
	}
	return
}
