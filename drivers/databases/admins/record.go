package admins

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/admins"
	"main.go/drivers/databases/jobs"
)

type Admins struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Email       string `gorm:"unique"`
	Address     string
	Password    string
	CompanyId   int
	JobsCreated []jobs.Jobs `gorm:"foreignKey:CreatedBy;references:Id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (admin *Admins) ToDomain() admins.Domain {
	return admins.Domain{
		Id:          admin.Id,
		Name:        admin.Name,
		Email:       admin.Email,
		Address:     admin.Address,
		CompanyId:   admin.CompanyId,
		JobsCreated: jobs.ListToDomain(admin.JobsCreated),
		Password:    admin.Password,
		CreatedAt:   admin.CreatedAt,
		UpdatedAt:   admin.UpdatedAt,
	}
}

func ListToDomain(admins []Admins) (result []admins.Domain) {
	for _, admin := range admins {
		result = append(result, admin.ToDomain())
	}
	return
}

func FromDomain(domain admins.Domain) Admins {
	return Admins{
		Id:          domain.Id,
		Name:        domain.Name,
		Email:       domain.Email,
		Address:     domain.Address,
		CompanyId:   domain.CompanyId,
		JobsCreated: jobs.ListFromDomain(domain.JobsCreated),
		Password:    domain.Password,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func ListFromDomain(admins []admins.Domain) (result []Admins) {
	for _, admin := range admins {
		result = append(result, FromDomain(admin))
	}
	return
}
