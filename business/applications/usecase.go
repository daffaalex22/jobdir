package applications

import (
	"context"
	"time"
)

type ApplicationUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewApplicationUsecase(repo Repository, timeout time.Duration) Usecase {
	return &ApplicationUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// func (uc *ApplicationUsecase) FillApplications(ctx context.Context, categories []CategoryDomain) ([]CategoryDomain, error) {

// 	res, err := uc.Repo.FillApplications(ctx, categories)
// 	if err != nil {
// 		return categories, err
// 	}

// 	return res, nil
// }

func (uc *ApplicationUsecase) CreateApplication(ctx context.Context, domain Domain) (Domain, error) {

	// if domain.Title == "" {
	// 	return Domain{}, errors.New("title empty")
	// }

	application, err := uc.Repo.CreateApplication(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return application, nil
}

func (uc *ApplicationUsecase) GetAllApplications(c context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	Applications, err := uc.Repo.GetAllApplications(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return Applications, nil
}

// func (uc *ApplicationUsecase) GetApplicationById(c context.Context, id int) (Domain, error) {
// 	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
// 	defer cancel()

// 	Application, err := uc.Repo.GetApplicationById(ctx, id)
// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return Application, nil
// }

func (uc *ApplicationUsecase) DeleteAllApplications(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	err := uc.Repo.DeleteAllApplications(ctx)
	if err != nil {
		return err
	}

	return nil
}

// func (uc *ApplicationUsecase) DeleteApplicationById(c context.Context, id int) (Domain, error) {
// 	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
// 	defer cancel()

// 	Application, err := uc.Repo.DeleteApplicationById(ctx, id)
// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return Application, nil
// }

// func (uc *ApplicationUsecase) SearchApplications(c context.Context, title string) ([]Domain, error) {
// 	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
// 	defer cancel()

// 	Application, err := uc.Repo.SearchApplications(ctx, title)
// 	if err != nil {
// 		return []Domain{}, err
// 	}

// 	return Application, nil
// }

// func (uc *ApplicationUsecase) FilterApplicationByCategory(c context.Context, categoryId int) ([]Domain, error) {
// 	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
// 	defer cancel()

// 	Application, err := uc.Repo.FilterApplicationByCategory(ctx, categoryId)
// 	if err != nil {
// 		return []Domain{}, err
// 	}

// 	return Application, nil
// }
