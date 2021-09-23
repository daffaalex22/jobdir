package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	Email     string
	Address   string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
}
