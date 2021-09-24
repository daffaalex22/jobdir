package jobs

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Title     string
	Category  string
	JobDesc   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetJobById(ctx context.Context, id int) (Domain, error)
	CreateJob(ctx context.Context, domain Domain) (Domain, error)
	DeleteAllJobs(ctx context.Context) error
	DeleteJobById(ctx context.Context, id int) (Domain, error)
	GetAllJobs(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	GetJobById(ctx context.Context, id int) (Domain, error)
	CreateJob(ctx context.Context, domain Domain) (Domain, error)
	DeleteAllJobs(ctx context.Context) error
	DeleteJobById(ctx context.Context, id int) (Domain, error)
	GetAllJobs(ctx context.Context) ([]Domain, error)
}
