package users

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/users"
)

type Users struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Address   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Password,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Password,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
