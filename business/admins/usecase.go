package admins

import (
	"context"
	"errors"
	"time"

	"main.go/app/middlewares"
)

type AdminUsecase struct {
	ConfigJWT      *middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewAdminUsecase(repo Repository, timeout time.Duration, configJWT *middlewares.ConfigJWT) Usecase {
	return &AdminUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *AdminUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}

	admin, err := uc.Repo.Login(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	admin.Token, err = uc.ConfigJWT.GenerateTokenJWT(admin.Id)
	if err != nil {
		return Domain{}, err
	}
	return admin, nil
}

func (uc *AdminUsecase) GetAdminById(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	admin, err := uc.Repo.GetAdminById(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return admin, nil
}

func (uc *AdminUsecase) GetAllAdmin(c context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	admin, err := uc.Repo.GetAllAdmin(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return admin, nil
}

func (uc *AdminUsecase) UpdateAdmin(c context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	// domain.UpdatedAt = time.Now()
	admin, err := uc.Repo.UpdateAdmin(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return admin, nil
}

func (uc *AdminUsecase) DeleteAdmin(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	admin, err := uc.Repo.DeleteAdmin(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return admin, nil
}

func (uc *AdminUsecase) RegisterAdmin(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}

	Admin, err := uc.Repo.RegisterAdmin(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return Admin, nil
}

func (uc *AdminUsecase) HardDeleteAllAdmins(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	err := uc.Repo.HardDeleteAllAdmins(ctx)
	if err != nil {
		return err
	}

	return nil
}
