package requests

import (
	"main.go/business/jobs"
)

type JobCreate struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	JobDesc  string `json:"jobDesc"`
}

func (job *JobCreate) ToDomain() jobs.Domain {
	return jobs.Domain{
		Title:    job.Title,
		Category: job.Category,
		JobDesc:  job.JobDesc,
	}
}
