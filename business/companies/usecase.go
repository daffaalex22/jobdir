package companies

import (
	"context"
	"time"
)

type CompanyUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCompanyUsecase(repo Repository, timeout time.Duration) Usecase {
	return &CompanyUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *CompanyUsecase) GetCompanyById(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	company, err := uc.Repo.GetCompanyById(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return company, nil
}

func (uc *CompanyUsecase) GetAllCompany(c context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	company, err := uc.Repo.GetAllCompany(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return company, nil
}

func (uc *CompanyUsecase) UpdateCompany(c context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	company, err := uc.Repo.UpdateCompany(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return company, nil
}

func (uc *CompanyUsecase) DeleteCompany(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	company, err := uc.Repo.DeleteCompany(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return company, nil
}

func (uc *CompanyUsecase) RegisterCompany(ctx context.Context, domain Domain) (Domain, error) {

	company, err := uc.Repo.RegisterCompany(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return company, nil
}
