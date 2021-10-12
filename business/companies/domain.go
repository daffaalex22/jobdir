package companies

import (
	"context"
	"time"

	"main.go/business/admins"
	"main.go/business/jobs"
)

type Domain struct {
	Id           int
	Name         string
	Address      string
	Description  string
	IsTopCompany bool
	Admins       []admins.Domain
	Jobs         []jobs.Domain
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	RegisterCompany(ctx context.Context, domain Domain) (Domain, error)
	GetCompanyById(ctx context.Context, id int) (Domain, error)
	GetAllCompany(ctx context.Context) ([]Domain, error)
	UpdateCompany(ctx context.Context, domain Domain) (Domain, error)
	DeleteCompany(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	RegisterCompany(ctx context.Context, domain Domain) (Domain, error)
	GetCompanyById(ctx context.Context, id int) (Domain, error)
	GetAllCompany(ctx context.Context) ([]Domain, error)
	UpdateCompany(ctx context.Context, domain Domain) (Domain, error)
	DeleteCompany(ctx context.Context, id int) (Domain, error)
}
