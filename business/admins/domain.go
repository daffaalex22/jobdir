package admins

import (
	"context"
	"time"
)

type Domain struct {
	Id          int
	Name        string
	Email       string
	Address     string
	CompanyName string
	Password    string
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	RegisterAdmin(ctx context.Context, domain Domain) (Domain, error)
	GetAdminById(ctx context.Context, id int) (Domain, error)
	GetAllAdmin(ctx context.Context) ([]Domain, error)
	UpdateAdmin(ctx context.Context, domain Domain) (Domain, error)
	DeleteAdmin(ctx context.Context, id int) (Domain, error)
	HardDeleteAllAdmins(ctx context.Context) error
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	RegisterAdmin(ctx context.Context, domain Domain) (Domain, error)
	GetAdminById(ctx context.Context, id int) (Domain, error)
	GetAllAdmin(ctx context.Context) ([]Domain, error)
	UpdateAdmin(ctx context.Context, domain Domain) (Domain, error)
	DeleteAdmin(ctx context.Context, id int) (Domain, error)
	HardDeleteAllAdmins(ctx context.Context) error
}
