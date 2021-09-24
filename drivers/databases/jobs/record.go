package jobs

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/jobs"
)

type Jobs struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Category  string
	JobDesc   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (job *Jobs) ToDomain() jobs.Domain {
	return jobs.Domain{
		Id:        job.Id,
		Title:     job.Title,
		Category:  job.Category,
		JobDesc:   job.JobDesc,
		CreatedAt: job.CreatedAt,
		UpdatedAt: job.UpdatedAt,
	}
}

func ListToDomain(jobs []Jobs) (result []jobs.Domain) {
	for _, job := range jobs {
		result = append(result, job.ToDomain())
	}
	return
}

func FromDomain(domain jobs.Domain) Jobs {
	return Jobs{
		Id:        domain.Id,
		Title:     domain.Title,
		Category:  domain.Category,
		JobDesc:   domain.JobDesc,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
