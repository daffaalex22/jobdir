package companies_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/business/companies"
	_mockcompanyRepository "main.go/business/companies/mocks"
)

var companyRepository _mockcompanyRepository.Repository
var companyService companies.Usecase
var companyDomain companies.Domain
var companiesDomain []companies.Domain

func setup() {
	companyService = companies.NewCompanyUsecase(&companyRepository, time.Hour*1 /*, configJWT */)
	companyDomain = companies.Domain{
		Id:           1,
		Name:         "Gojex",
		Address:      "Bandungs",
		Description:  "Nadiem Makarim utk Presiden",
		IsTopCompany: true,
	}
	companiesDomain = append(companiesDomain, companyDomain)
}

func TestGetCompanyById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		companyRepository.On("GetCompanyById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(companyDomain, nil).Once()

		Company, err := companyService.GetCompanyById(context.Background(), companyDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, Company)

		// companyRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		companyRepository.On("GetCompanyById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(companies.Domain{}, errors.New(mock.Anything)).Once()

		Company, err := companyService.GetCompanyById(context.Background(), companyDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, Company, companies.Domain{})

		// companyRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	companyRepository.On("GetCompanyById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(companyDomain, nil).Once()

	// 	Company, err := companyService.GetCompanyById(context.Background(), companyDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, companies.Domain{}, Company)

	// companyRepository.AssertExpectations(t)
	// })
}

func TestGetAllCompany(t *testing.T) {
	setup()
	// companiesDomain = append(companiesDomain, companyDomain)

	t.Run("Test case 1", func(t *testing.T) {
		companyRepository.On("GetAllCompany",
			mock.Anything).Return(companiesDomain, nil).Once()

		Company, err := companyService.GetAllCompany(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, Company)

		// companyRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		companyRepository.On("GetAllCompany",
			mock.Anything).Return([]companies.Domain{}, errors.New(mock.Anything)).Once()

		Company, err := companyService.GetAllCompany(context.Background())

		assert.Error(t, err)
		assert.Equal(t, Company, []companies.Domain{})

		// companyRepository.AssertExpectations(t)
	})
}

func TestDeleteCompany(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		companyRepository.On("DeleteCompany",
			mock.Anything,
			mock.AnythingOfType("int")).Return(companyDomain, nil).Once()

		Company, err := companyService.DeleteCompany(context.Background(), companyDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, Company)

		// companyRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Return Error", func(t *testing.T) {
		companyRepository.On("DeleteCompany",
			mock.Anything,
			mock.AnythingOfType("int")).Return(companies.Domain{}, errors.New(mock.Anything)).Once()

		Company, err := companyService.DeleteCompany(context.Background(), companyDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, Company, companies.Domain{})

		// companyRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	companyRepository.On("GetCompanyById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(companyDomain, nil).Once()

	// 	Company, err := companyService.GetCompanyById(context.Background(), companyDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, companies.Domain{}, Company)

	// companyRepository.AssertExpectations(t)
	// })
}

func TestRegisterCompany(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Create", func(t *testing.T) {
		companyRepository.On("RegisterCompany",
			mock.Anything,
			mock.AnythingOfType("companies.Domain")).Return(companyDomain, nil).Once()
		Company, err := companyService.RegisterCompany(context.Background(), companies.Domain{
			Name:         "Gojex",
			Address:      "Bandungs",
			Description:  "Nadiem Makarim utk Presiden",
			IsTopCompany: true,
		})

		assert.Nil(t, err)
		assert.NotNil(t, Company)
	})

	t.Run("Test Case 2 | Return Error", func(t *testing.T) {
		companyRepository.On("RegisterCompany",
			mock.Anything,
			mock.AnythingOfType("companies.Domain")).Return(companies.Domain{}, errors.New(mock.Anything)).Once()
		Company, err := companyService.RegisterCompany(context.Background(), companies.Domain{
			Name:         "Gojex",
			Address:      "Bandungs",
			Description:  "Nadiem Makarim utk Presiden",
			IsTopCompany: true,
		})
		assert.Error(t, err)
		assert.Equal(t, Company, companies.Domain{})
	})

	// t.Run("Test Case 3 | Invalid Company Empty", func(t *testing.T) {
	// 	companyRepository.On("RegisterCompany",
	// 		mock.Anything,
	// 		mock.AnythingOfType("companies.Domain")).Return(companyDomain, nil).Once()
	// 	_, err := companyService.RegisterCompany(context.Background(), companies.Domain{
	// 		Company: "",
	// 	})
	// 	assert.NotNil(t, err)
	// })

}

func TestHardDeleteAllCompany(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		companyRepository.On("HardDeleteAllCompanies",
			mock.Anything).Return(nil).Once()

		err := companyService.HardDeleteAllCompanies(context.Background())

		assert.NoError(t, err)

		// companyRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		companyRepository.On("HardDeleteAllCompanies",
			mock.Anything).Return(errors.New(mock.Anything)).Once()

		err := companyService.HardDeleteAllCompanies(context.Background())

		assert.Error(t, err)

		// companyRepository.AssertExpectations(t)
	})
}

func TestUpdateCompany(t *testing.T) {
	setup()
	// adminsDomain = append(adminsDomain, companyDomain)

	t.Run("Test case 1", func(t *testing.T) {
		companyRepository.On("UpdateCompany",
			mock.Anything,
			mock.AnythingOfType("companies.Domain")).Return(companyDomain, nil).Once()

		admin, err := companyService.UpdateCompany(context.Background(), companies.Domain{
			Name:         "Pablo",
			Address:      "Bandungs",
			Description:  "Kemantapan Perusahaan",
			IsTopCompany: false,
		})

		assert.NoError(t, err)
		assert.Equal(t, admin.Address, "Bandungs")

		// companyRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		companyRepository.On("UpdateCompany",
			mock.Anything,
			mock.AnythingOfType("companies.Domain")).Return(companies.Domain{}, errors.New("Unexpected Error")).Once()

		admin, err := companyService.UpdateCompany(context.Background(), companies.Domain{
			Name:         "Pablo",
			Address:      "Bandungs",
			Description:  "Kemantapan Perusahaan",
			IsTopCompany: false,
		})

		assert.Error(t, err)
		assert.Equal(t, admin, companies.Domain{})

		// companyRepository.AssertExpectations(t)
	})
}
