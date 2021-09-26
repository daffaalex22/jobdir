package admins

import (
	"context"
	"errors"
	"time"

	"github.com/spf13/viper"
	"main.go/app/middlewares"
)

type AdminUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewAdminUsecase(repo Repository, timeout time.Duration) Usecase {
	return &AdminUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *AdminUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {

	if email == "" {
		return Domain{}, errors.New("email empty")
	}

	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	Admin, err := uc.Repo.Login(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}
	JWTConf := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Admin.Token, err = JWTConf.GenerateTokenJWT(Admin.Id)
	if err != nil {
		return Domain{}, err
	}
	return Admin, nil
}

func (uc *AdminUsecase) GetAdminById(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	Admin, err := uc.Repo.GetAdminById(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return Admin, nil
}

func (uc *AdminUsecase) GetAllAdmin(c context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	Admin, err := uc.Repo.GetAllAdmin(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return Admin, nil
}

func (uc *AdminUsecase) UpdateAdmin(c context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	// domain.UpdatedAt = time.Now()
	Admin, err := uc.Repo.UpdateAdmin(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return Admin, nil
}

func (uc *AdminUsecase) DeleteAdmin(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	Admin, err := uc.Repo.DeleteAdmin(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return Admin, nil
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
