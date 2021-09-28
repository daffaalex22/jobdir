package responses

import (
	"time"

	"main.go/business/jobs"
	_appResponses "main.go/controllers/applications/responses"
	_catResponses "main.go/controllers/categories/responses"
)

type JobResponse struct {
	Id           int                                 `json:"id"`
	Title        string                              `json:"title"`
	CategoryId   int                                 `json:"categoryId"`
	Category     _catResponses.CategoryResponse      `json:"category"`
	JobDesc      string                              `json:"jobDesc"`
	CreatedBy    int                                 `json:"createdBy"`
	CompanyId    int                                 `json:"companyId"`
	Applications []_appResponses.ApplicationResponse `json:"applications"`
	CreatedAt    time.Time                           `json:"createdAt"`
	UpdatedAt    time.Time                           `json:"updatedAt"`
}

func FromDomain(domain jobs.Domain) JobResponse {
	return JobResponse{
		Id:           domain.Id,
		Title:        domain.Title,
		CategoryId:   domain.CategoryId,
		Category:     _catResponses.FromDomain(domain.Category),
		JobDesc:      domain.JobDesc,
		CreatedBy:    domain.CreatedBy,
		CompanyId:    domain.CompanyId,
		Applications: _appResponses.ListFromDomain(domain.Applications),
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
