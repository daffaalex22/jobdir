package jobs

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/jobs"
	"main.go/drivers/databases/applications"
	"main.go/drivers/databases/categories"
)

type Jobs struct {
	Id         int `gorm:"primaryKey"`
	Title      string
	CategoryId int
	Category   categories.Categories `gorm:"foreignKey:CategoryId"`
	JobDesc    string
	CreatedBy  int
	CompanyId  int
	// Company      companies.Companies         `gorm:"foreignKey:CompanyId"`
	Applications []applications.Applications `gorm:"foreignKey:JobId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (job *Jobs) ToDomain() jobs.Domain {
	return jobs.Domain{
		Id:         job.Id,
		Title:      job.Title,
		CategoryId: job.CategoryId,
		Category:   job.Category.ToDomain(),
		JobDesc:    job.JobDesc,
		CreatedBy:  job.CreatedBy,
		CompanyId:  job.CompanyId,
		// Company:      job.Company.ToDomain(),
		Applications: applications.ListToDomain(job.Applications),
		CreatedAt:    job.CreatedAt,
		UpdatedAt:    job.UpdatedAt,
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
		Category:   categories.FromDomain(domain.Category),
		JobDesc:    domain.JobDesc,
		CreatedBy:  domain.CreatedBy,
		CompanyId:  domain.CompanyId,
		// Company:      companies.FromDomain(domain.Company),
		Applications: applications.ListFromDomain(domain.Applications),
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func ListFromDomain(jobs []jobs.Domain) (result []Jobs) {
	for _, job := range jobs {
		result = append(result, FromDomain(job))
	}
	return
}
