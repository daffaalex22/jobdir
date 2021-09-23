package users

import (
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {

	if email == "" {
		return Domain{}, errors.New("email empty")
	}

	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	user, err := uc.Repo.Login(ctx, email, password)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUsecase) GetUserById(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	user, err := uc.Repo.GetUserById(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUsecase) GetAllUser(c context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	user, err := uc.Repo.GetAllUser(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *UserUsecase) UpdateUser(c context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	// domain.UpdatedAt = time.Now()
	user, err := uc.Repo.UpdateUser(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUsecase) DeleteUser(c context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	user, err := uc.Repo.DeleteUser(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
