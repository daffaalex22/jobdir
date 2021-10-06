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

// func (uc *CompanyUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {

// 	if email == "" {
// 		return Domain{}, errors.New("email empty")
// 	}

// 	if password == "" {
// 		return Domain{}, errors.New("password empty")
// 	}

// 	Company, err := uc.Repo.Login(ctx, email, password)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	JWTConf := middlewares.ConfigJWT{
// 		SecretJWT:       viper.GetString(`jwt.secret`),
// 		ExpiresDuration: viper.GetInt(`jwt.expired`),
// 	}

// 	Company.Token, err = JWTConf.GenerateTokenJWT(Company.Id)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	return Company, nil
// }

func (uc *CompanyUsecase) GetCompanyById(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	Company, err := uc.Repo.GetCompanyById(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return Company, nil
}

func (uc *CompanyUsecase) GetAllCompany(c context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	Company, err := uc.Repo.GetAllCompany(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return Company, nil
}

func (uc *CompanyUsecase) UpdateCompany(c context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	// domain.UpdatedAt = time.Now()
	Company, err := uc.Repo.UpdateCompany(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return Company, nil
}

func (uc *CompanyUsecase) DeleteCompany(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	Company, err := uc.Repo.DeleteCompany(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return Company, nil
}

func (uc *CompanyUsecase) RegisterCompany(ctx context.Context, domain Domain) (Domain, error) {

	company, err := uc.Repo.RegisterCompany(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return company, nil
}

func (uc *CompanyUsecase) HardDeleteAllCompanies(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	err := uc.Repo.HardDeleteAllCompanies(ctx)
	if err != nil {
		return err
	}

	return nil
}
