package jobs

import (
	"context"
	"time"

	"main.go/business/applications"
	"main.go/business/categories"
)

type Domain struct {
	Id           int
	Title        string
	CategoryId   int
	Category     categories.Domain
	JobDesc      string
	CreatedBy    int
	CompanyId    int
	Applications []applications.Domain
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// type CategoryDomain struct {
// 	Id       int
// 	Category string
// 	// CategoryDescription string
// 	Jobs      []Domain
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type Usecase interface {
	GetJobById(ctx context.Context, id int) (Domain, error)
	CreateJob(ctx context.Context, domain Domain) (Domain, error)
	DeleteAllJobs(ctx context.Context) error
	DeleteJobById(ctx context.Context, id int) (Domain, error)
	GetAllJobs(ctx context.Context) ([]Domain, error)
	SearchJobs(ctx context.Context, title string) ([]Domain, error)
	FilterJobByCategory(ctx context.Context, categoryId int) ([]Domain, error)
	// FillJobs(ctx context.Context, categories []CategoryDomain) ([]CategoryDomain, error)
}

type Repository interface {
	GetJobById(ctx context.Context, id int) (Domain, error)
	CreateJob(ctx context.Context, domain Domain) (Domain, error)
	DeleteAllJobs(ctx context.Context) error
	DeleteJobById(ctx context.Context, id int) (Domain, error)
	GetAllJobs(ctx context.Context) ([]Domain, error)
	SearchJobs(ctx context.Context, title string) ([]Domain, error)
	FilterJobByCategory(ctx context.Context, categoryId int) ([]Domain, error)
	// FillJobs(ctx context.Context, categories []CategoryDomain) ([]CategoryDomain, error)
}
