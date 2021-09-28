package requests

import (
	"main.go/business/jobs"
)

type JobCreate struct {
	Title      string `json:"title"`
	CategoryId int    `json:"categoryId"`
	CompanyId  int    `json:"companyId"`
	JobDesc    string `json:"jobDesc"`
	CreatedBy  int    `json:"createdBy"`
}

func (job *JobCreate) ToDomain() jobs.Domain {
	return jobs.Domain{
		Title:      job.Title,
		CategoryId: job.CategoryId,
		CompanyId:  job.CompanyId,
		JobDesc:    job.JobDesc,
		CreatedBy:  job.CreatedBy,
	}
}

func ListToDomain(jobs []JobCreate) (result []jobs.Domain) {
	for _, job := range jobs {
		result = append(result, job.ToDomain())
	}
	return
}
