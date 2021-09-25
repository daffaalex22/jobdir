package requests

import (
	"main.go/business/jobs"
)

type JobCreate struct {
	Title      string `json:"title"`
	CategoryId int    `json:"categoryId"`
	JobDesc    string `json:"jobDesc"`
}

func (job *JobCreate) ToDomain() jobs.Domain {
	return jobs.Domain{
		Title:      job.Title,
		CategoryId: job.CategoryId,
		JobDesc:    job.JobDesc,
	}
}

func ListToDomain(jobs []JobCreate) (result []jobs.Domain) {
	for _, job := range jobs {
		result = append(result, job.ToDomain())
	}
	return
}
