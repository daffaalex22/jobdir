package applications

import (
	"context"
	"time"
)

type Domain struct {
	UserId    int
	JobId     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type CategoryDomain struct {
// 	Id       int
// 	Category string
// 	// CategoryDescription string
// 	Application      []Domain
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type Usecase interface {
	// GetApplicationById(ctx context.Context, id int) (Domain, error)
	CreateApplication(ctx context.Context, domain Domain) (Domain, error)
	DeleteAllApplications(ctx context.Context) error
	// DeleteApplicationById(ctx context.Context, id int) (Domain, error)
	GetAllApplications(ctx context.Context) ([]Domain, error)
	// SearchApplication(ctx context.Context, title string) ([]Domain, error)
	// FilterApplicationByCategory(ctx context.Context, categoryId int) ([]Domain, error)
	// FillApplication(ctx context.Context, categories []CategoryDomain) ([]CategoryDomain, error)
}

type Repository interface {
	// GetApplicationById(ctx context.Context, id int) (Domain, error)
	CreateApplication(ctx context.Context, domain Domain) (Domain, error)
	DeleteAllApplications(ctx context.Context) error
	// DeleteApplicationById(ctx context.Context, id int) (Domain, error)
	GetAllApplications(ctx context.Context) ([]Domain, error)
	// SearchApplication(ctx context.Context, title string) ([]Domain, error)
	// FilterApplicationByCategory(ctx context.Context, categoryId int) ([]Domain, error)
	// FillApplication(ctx context.Context, categories []CategoryDomain) ([]CategoryDomain, error)
}
