package companies_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/business/companies"
	_mockCompanyRepository "main.go/business/companies/mocks"
)

var CompanyRepository _mockCompanyRepository.Repository
var companieService companies.Usecase
var CompanyDomain companies.Domain
var companiesDomain []companies.Domain

// var configJWT *middlewares.ConfigJWT

func setup() {
	companieService = companies.NewCompanyUsecase(&CompanyRepository, time.Hour*1 /*, configJWT */)
	CompanyDomain = companies.Domain{
		Id:           1,
		Name:         "Gojex",
		Address:      "Bandungs",
		Description:  "Nadiem Makarim utk Presiden",
		IsTopCompany: true,
	}
	companiesDomain = append(companiesDomain, CompanyDomain)
}

func TestGetCompanyById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		CompanyRepository.On("GetCompanyById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(CompanyDomain, nil).Once()

		Company, err := companieService.GetCompanyById(context.Background(), CompanyDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, Company)

		CompanyRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		CompanyRepository.On("GetCompanyById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(companies.Domain{}, errors.New(mock.Anything)).Once()

		Company, err := companieService.GetCompanyById(context.Background(), CompanyDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, Company, companies.Domain{})

		CompanyRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	CompanyRepository.On("GetCompanyById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(CompanyDomain, nil).Once()

	// 	Company, err := companieService.GetCompanyById(context.Background(), CompanyDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, companies.Domain{}, Company)

	// 	CompanyRepository.AssertExpectations(t)
	// })
}

func TestGetAllCompany(t *testing.T) {
	setup()
	// companiesDomain = append(companiesDomain, CompanyDomain)

	t.Run("Test case 1", func(t *testing.T) {
		CompanyRepository.On("GetAllCompany",
			mock.Anything).Return(companiesDomain, nil).Once()

		Company, err := companieService.GetAllCompany(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, Company)

		CompanyRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		CompanyRepository.On("GetAllCompany",
			mock.Anything).Return([]companies.Domain{}, errors.New(mock.Anything)).Once()

		Company, err := companieService.GetAllCompany(context.Background())

		assert.Error(t, err)
		assert.Equal(t, Company, []companies.Domain{})

		CompanyRepository.AssertExpectations(t)
	})
}

func TestDeleteCompany(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		CompanyRepository.On("DeleteCompany",
			mock.Anything,
			mock.AnythingOfType("int")).Return(CompanyDomain, nil).Once()

		Company, err := companieService.DeleteCompany(context.Background(), CompanyDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, Company)

		CompanyRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Return Error", func(t *testing.T) {
		CompanyRepository.On("DeleteCompany",
			mock.Anything,
			mock.AnythingOfType("int")).Return(companies.Domain{}, errors.New(mock.Anything)).Once()

		Company, err := companieService.DeleteCompany(context.Background(), CompanyDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, Company, companies.Domain{})

		CompanyRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	CompanyRepository.On("GetCompanyById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(CompanyDomain, nil).Once()

	// 	Company, err := companieService.GetCompanyById(context.Background(), CompanyDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, companies.Domain{}, Company)

	// 	CompanyRepository.AssertExpectations(t)
	// })
}

func TestRegisterCompany(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Create", func(t *testing.T) {
		CompanyRepository.On("RegisterCompany",
			mock.Anything,
			mock.AnythingOfType("companies.Domain")).Return(CompanyDomain, nil).Once()
		Company, err := companieService.RegisterCompany(context.Background(), companies.Domain{
			Name:         "Gojex",
			Address:      "Bandungs",
			Description:  "Nadiem Makarim utk Presiden",
			IsTopCompany: true,
		})

		assert.Nil(t, err)
		assert.NotNil(t, Company)
	})

	t.Run("Test Case 2 | Return Error", func(t *testing.T) {
		CompanyRepository.On("RegisterCompany",
			mock.Anything,
			mock.AnythingOfType("companies.Domain")).Return(companies.Domain{}, errors.New(mock.Anything)).Once()
		Company, err := companieService.RegisterCompany(context.Background(), companies.Domain{
			Name:         "Gojex",
			Address:      "Bandungs",
			Description:  "Nadiem Makarim utk Presiden",
			IsTopCompany: true,
		})
		assert.Error(t, err)
		assert.Equal(t, Company, companies.Domain{})
	})

	// t.Run("Test Case 3 | Invalid Company Empty", func(t *testing.T) {
	// 	CompanyRepository.On("RegisterCompany",
	// 		mock.Anything,
	// 		mock.AnythingOfType("companies.Domain")).Return(CompanyDomain, nil).Once()
	// 	_, err := companieService.RegisterCompany(context.Background(), companies.Domain{
	// 		Company: "",
	// 	})
	// 	assert.NotNil(t, err)
	// })

}

func TestHardDeleteAllCompany(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		CompanyRepository.On("HardDeleteAllCompanies",
			mock.Anything).Return(nil).Once()

		err := companieService.HardDeleteAllCompanies(context.Background())

		assert.NoError(t, err)

		// CompanyRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		CompanyRepository.On("HardDeleteAllCompanies",
			mock.Anything).Return(errors.New(mock.Anything)).Once()

		err := companieService.HardDeleteAllCompanies(context.Background())

		assert.Error(t, err)

		// CompanyRepository.AssertExpectations(t)
	})
}
