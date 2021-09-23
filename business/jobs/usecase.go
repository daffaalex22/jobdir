package jobs

import (
	"time"
)

type JobUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewJobUsecase(repo Repository, timeout time.Duration) Usecase {
	return &JobUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// func (uc *JobUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {

// 	if email == "" {
// 		return Domain{}, errors.New("email empty")
// 	}

// 	if password == "" {
// 		return Domain{}, errors.New("password empty")
// 	}

// 	user, err := uc.Repo.Login(ctx, email, password)

// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return user, nil
// }
