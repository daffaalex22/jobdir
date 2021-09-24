package jobs

import (
	"context"
	"errors"
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

func (uc *JobUsecase) CreateJob(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Title == "" {
		return Domain{}, errors.New("title empty")
	}

	job, err := uc.Repo.CreateJob(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return job, nil
}

func (uc *JobUsecase) GetAllJobs(c context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	jobs, err := uc.Repo.GetAllJobs(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return jobs, nil
}

func (uc *JobUsecase) GetJobById(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	job, err := uc.Repo.GetJobById(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return job, nil
}

func (uc *JobUsecase) DeleteAllJobs(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	err := uc.Repo.DeleteAllJobs(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (uc *JobUsecase) DeleteJobById(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	Job, err := uc.Repo.DeleteJobById(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return Job, nil
}
