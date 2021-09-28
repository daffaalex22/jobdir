package categories

import (
	"context"
	"errors"
	"time"
)

type CategoryUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCategoryUsecase(repo Repository, timeout time.Duration) Usecase {
	return &CategoryUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *CategoryUsecase) CreateCategory(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Category == "" {
		return Domain{}, errors.New("category empty")
	}

	category, err := uc.Repo.CreateCategory(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return category, nil
}

func (uc *CategoryUsecase) GetCategoryById(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	category, err := uc.Repo.GetCategoryById(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return category, nil
}

func (uc *CategoryUsecase) GetAllCategory(c context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	category, err := uc.Repo.GetAllCategory(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return category, nil
}

func (uc *CategoryUsecase) DeleteCategoryById(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	category, err := uc.Repo.DeleteCategoryById(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return category, nil
}
