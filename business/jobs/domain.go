package jobs

import (
	"time"
)

type Domain struct {
	Id        int
	Title     string
	Category  string
	JobDesc   string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
}

type Repository interface {
}
