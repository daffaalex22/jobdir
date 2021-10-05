package applications

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/applications"
)

type Applications struct {
	UserId    int `gorm:"primaryKey;autoIncrement:false"`
	JobId     int `gorm:"primaryKey:autoIncrement:false"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (application *Applications) ToDomain() applications.Domain {
	return applications.Domain{
		UserId:    application.UserId,
		JobId:     application.JobId,
		Status:    application.Status,
		CreatedAt: application.CreatedAt,
		UpdatedAt: application.UpdatedAt,
	}
}

func ListToDomain(Applications []Applications) (result []applications.Domain) {
	for _, application := range Applications {
		result = append(result, application.ToDomain())
	}
	return
}

func FromDomain(domain applications.Domain) Applications {
	return Applications{
		UserId:    domain.UserId,
		JobId:     domain.JobId,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ListFromDomain(Applications []applications.Domain) (result []Applications) {
	for _, job := range Applications {
		result = append(result, FromDomain(job))
	}
	return
}
