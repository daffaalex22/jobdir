package responses

import (
	"time"

	"main.go/business/jobs"
)

type JobResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Category  string    `json:"category"`
	JobDesc   string    `json:"jobDesc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain jobs.Domain) JobResponse {
	return JobResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		Category:  domain.Category,
		JobDesc:   domain.JobDesc,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ListFromDomain(domain []jobs.Domain) (response []JobResponse) {
	for _, job := range domain {
		response = append(response, FromDomain(job))
	}
	return
}
