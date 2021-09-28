package categories

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	CreateCategory(ctx context.Context, domain Domain) (Domain, error)
	GetCategoryById(ctx context.Context, id int) (Domain, error)
	GetAllCategory(ctx context.Context) ([]Domain, error)
	DeleteCategoryById(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	CreateCategory(ctx context.Context, domain Domain) (Domain, error)
	GetCategoryById(ctx context.Context, id int) (Domain, error)
	GetAllCategory(ctx context.Context) ([]Domain, error)
	DeleteCategoryById(ctx context.Context, id int) (Domain, error)
}
