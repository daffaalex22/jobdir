package users

import (
	"context"
	"time"

	"main.go/business/applications"
)

type Domain struct {
	Id           int
	Name         string
	Email        string
	Address      string
	Applications []applications.Domain
	Password     string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	RegisterUser(ctx context.Context, domain Domain) (Domain, error)
	GetUserById(ctx context.Context, id int) (Domain, error)
	GetAllUser(ctx context.Context) ([]Domain, error)
	UpdateUser(ctx context.Context, domain Domain) (Domain, error)
	DeleteUser(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	RegisterUser(ctx context.Context, domain Domain) (Domain, error)
	GetUserById(ctx context.Context, id int) (Domain, error)
	GetAllUser(ctx context.Context) ([]Domain, error)
	UpdateUser(ctx context.Context, domain Domain) (Domain, error)
	DeleteUser(ctx context.Context, id int) (Domain, error)
}
