package users

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/users"
	"main.go/drivers/databases/applications"
)

type Users struct {
	Id           int `gorm:"primaryKey"`
	Name         string
	Email        string `gorm:"unique"`
	Address      string
	Password     string
	Applications []applications.Applications `gorm:"foreignKey:UserId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email,
		Address:      user.Address,
		Password:     user.Password,
		Applications: applications.ListToDomain(user.Applications),
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func ListToDomain(users []Users) (result []users.Domain) {
	for _, user := range users {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:           domain.Id,
		Name:         domain.Name,
		Email:        domain.Email,
		Address:      domain.Address,
		Applications: applications.ListFromDomain(domain.Applications),
		Password:     domain.Password,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
