package admins

import (
	"context"
	"time"

	"main.go/business/jobs"
)

type Domain struct {
	Id          int
	Name        string
	Email       string
	Address     string
	CompanyId   int
	JobsCreated []jobs.Domain
	Password    string
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	RegisterAdmin(ctx context.Context, domain Domain) (Domain, error)
	GetAdminById(ctx context.Context, id int) (Domain, error)
	GetAllAdmin(ctx context.Context) ([]Domain, error)
	UpdateAdmin(ctx context.Context, domain Domain) (Domain, error)
	DeleteAdmin(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	RegisterAdmin(ctx context.Context, domain Domain) (Domain, error)
	GetAdminById(ctx context.Context, id int) (Domain, error)
	GetAllAdmin(ctx context.Context) ([]Domain, error)
	UpdateAdmin(ctx context.Context, domain Domain) (Domain, error)
	DeleteAdmin(ctx context.Context, id int) (Domain, error)
}
