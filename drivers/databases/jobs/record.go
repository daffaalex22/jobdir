package jobs

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/jobs"
)

type Jobs struct {
	Id         int `gorm:"primaryKey"`
	Title      string
	CategoryId int
	JobDesc    string
	CreatedBy  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (job *Jobs) ToDomain() jobs.Domain {
	return jobs.Domain{
		Id:         job.Id,
		Title:      job.Title,
		CategoryId: job.CategoryId,
		JobDesc:    job.JobDesc,
		CreatedBy:  job.CreatedBy,
		CreatedAt:  job.CreatedAt,
		UpdatedAt:  job.UpdatedAt,
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
		Id:         domain.Id,
		Title:      domain.Title,
		CategoryId: domain.CategoryId,
		JobDesc:    domain.JobDesc,
		CreatedBy:  domain.CreatedBy,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func ListFromDomain(jobs []jobs.Domain) (result []Jobs) {
	for _, job := range jobs {
		result = append(result, FromDomain(job))
	}
	return
}
