package admins

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/admins"
)

type Admins struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Email       string `gorm:"unique"`
	Address     string
	Password    string
	CompanyName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (admin *Admins) ToDomain() admins.Domain {
	return admins.Domain{
		Id:        admin.Id,
		Name:      admin.Name,
		Email:     admin.Email,
		Address:   admin.Address,
		Password:  admin.Password,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
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
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
