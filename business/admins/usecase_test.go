package admins_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/app/middlewares"
	"main.go/business/admins"
	_mockAdminRepository "main.go/business/admins/mocks"
	"main.go/business/jobs"
)

var adminRepository _mockAdminRepository.Repository
var adminService admins.Usecase
var adminDomain admins.Domain
var adminsDomain []admins.Domain

var configJWT middlewares.ConfigJWT

func setup() {
	configJWT = middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	adminService = admins.NewAdminUsecase(&adminRepository, time.Hour*1, &configJWT)
	adminDomain = admins.Domain{
		Id:        1,
		Name:      "Pabby",
		Email:     "daaffa@net.usc",
		Address:   "Belanda",
		CompanyId: 1,
		JobsCreated: []jobs.Domain{
			{
				Id:         1,
				Title:      "Software Engineer",
				CategoryId: 1,
				JobDesc:    "Kerja Lembur Bagai Kuda",
				CreatedBy:  1,
				CompanyId:  1,
			},
		},
		Password: "kecoak11",
		Token:    "123",
	}
	adminsDomain = append(adminsDomain, adminDomain)
}

func TestLogin(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Login", func(t *testing.T) {
		adminRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("admins.Domain")).Return(adminDomain, nil).Once()
		admin, err := adminService.Login(context.Background(), admins.Domain{
			Email:    "daaffa@net.usc",
			Password: "kecoak11",
		})

		assert.Nil(t, err)
		assert.Equal(t, "Pabby", admin.Name)
	})

	t.Run("Test case 2 | Error Login", func(t *testing.T) {
		adminRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("admins.Domain")).Return(admins.Domain{}, errors.New("Unexpected Error")).Once()
		admin, err := adminService.Login(context.Background(), admins.Domain{
			Email:    "daaffa@net.usc",
			Password: "kecoak11",
		})

		assert.NotNil(t, err)
		assert.Equal(t, admin, admins.Domain{})
	})

	t.Run("Test Case 3 | Invalid Email Empty", func(t *testing.T) {

		_, err := adminService.Login(context.Background(), admins.Domain{
			Email:    "",
			Password: "kecoak11",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Invalid Password Empty", func(t *testing.T) {

		_, err := adminService.Login(context.Background(), admins.Domain{
			Email:    "daaffa@net.usc",
			Password: "",
		})
		assert.NotNil(t, err)
	})

}

func TestGetAdminById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		adminRepository.On("GetAdminById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(adminDomain, nil).Once()

		admin, err := adminService.GetAdminById(context.Background(), adminDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, admin)

		// adminRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		adminRepository.On("GetAdminById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(admins.Domain{}, errors.New("Unexpected Error")).Once()

		admin, err := adminService.GetAdminById(context.Background(), adminDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, admin, admins.Domain{})

		// adminRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	adminRepository.On("GetAdminById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(adminDomain, nil).Once()

	// 	admin, err := adminService.GetAdminById(context.Background(), adminDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, admins.Domain{}, admin)

	// adminRepository.AssertExpectations(t)
	// })
}

func TestGetAllAdmin(t *testing.T) {
	setup()
	// adminsDomain = append(adminsDomain, adminDomain)

	t.Run("Test case 1", func(t *testing.T) {
		adminRepository.On("GetAllAdmin",
			mock.Anything).Return(adminsDomain, nil).Once()

		admin, err := adminService.GetAllAdmin(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, admin)

		// adminRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		adminRepository.On("GetAllAdmin",
			mock.Anything).Return([]admins.Domain{}, errors.New("Unexpected Error")).Once()

		admin, err := adminService.GetAllAdmin(context.Background())

		assert.Error(t, err)
		assert.Equal(t, admin, []admins.Domain{})

		// adminRepository.AssertExpectations(t)
	})
}

func TestDeleteAdmin(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		adminRepository.On("DeleteAdmin",
			mock.Anything,
			mock.AnythingOfType("int")).Return(adminDomain, nil).Once()

		admin, err := adminService.DeleteAdmin(context.Background(), adminDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, admin)

		// adminRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		adminRepository.On("DeleteAdmin",
			mock.Anything,
			mock.AnythingOfType("int")).Return(admins.Domain{}, errors.New("Unexpected Error")).Once()

		admin, err := adminService.DeleteAdmin(context.Background(), adminDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, admin, admins.Domain{})

		// adminRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	adminRepository.On("GetAdminById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(adminDomain, nil).Once()

	// 	admin, err := adminService.GetAdminById(context.Background(), adminDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, admins.Domain{}, admin)

	// adminRepository.AssertExpectations(t)
	// })
}

func TestRegisterAdmin(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Registry", func(t *testing.T) {
		adminRepository.On("RegisterAdmin",
			mock.Anything,
			mock.AnythingOfType("admins.Domain")).Return(adminDomain, nil).Once()
		admin, err := adminService.RegisterAdmin(context.Background(), admins.Domain{
			Email:    "daaffa@net.usc",
			Password: "kecoak11",
		})

		assert.Nil(t, err)
		assert.Equal(t, "Pabby", admin.Name)
	})

	t.Run("Test case 2 | Error Registry", func(t *testing.T) {
		adminRepository.On("RegisterAdmin",
			mock.Anything,
			mock.AnythingOfType("admins.Domain")).Return(admins.Domain{}, errors.New("Unexpected Error")).Once()
		admin, err := adminService.RegisterAdmin(context.Background(), admins.Domain{
			Email:    "daaffa@net.usc",
			Password: "kecoak11",
		})

		assert.Error(t, err)
		assert.Equal(t, admin, admins.Domain{})
	})

	t.Run("Test Case 3 | Invalid Email Empty", func(t *testing.T) {
		adminRepository.On("RegisterAdmin",
			mock.Anything,
			mock.AnythingOfType("admins.Domain")).Return(adminDomain, nil).Once()
		_, err := adminService.RegisterAdmin(context.Background(), admins.Domain{
			Email:    "",
			Password: "kecoak11",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Password Empty", func(t *testing.T) {
		adminRepository.On("RegisterAdmin",
			mock.Anything,
			mock.AnythingOfType("admins.Domain")).Return(adminDomain, nil).Once()
		_, err := adminService.RegisterAdmin(context.Background(), admins.Domain{
			Email:    "daaffa@net.usc",
			Password: "",
		})
		assert.NotNil(t, err)
	})

}

func TestHardDeleteAllAdmin(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		adminRepository.On("HardDeleteAllAdmins",
			mock.Anything).Return(nil).Once()

		err := adminService.HardDeleteAllAdmins(context.Background())

		assert.NoError(t, err)

		// adminRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		adminRepository.On("HardDeleteAllAdmins",
			mock.Anything).Return(errors.New("Unexpected Error")).Once()

		err := adminService.HardDeleteAllAdmins(context.Background())

		assert.Error(t, err)

		// adminRepository.AssertExpectations(t)
	})
}

func TestUpdateAdmin(t *testing.T) {
	setup()
	// adminsDomain = append(adminsDomain, adminDomain)

	t.Run("Test case 1", func(t *testing.T) {
		adminRepository.On("UpdateAdmin",
			mock.Anything,
			mock.AnythingOfType("admins.Domain")).Return(adminDomain, nil).Once()

		admin, err := adminService.UpdateAdmin(context.Background(), admins.Domain{
			Email:    "daaffa@net.usc",
			Password: "kecoak11",
			Name:     "Pablo",
			Address:  "Belanda",
		})

		assert.NoError(t, err)
		assert.Equal(t, admin.Address, "Belanda")

		// adminRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		adminRepository.On("UpdateAdmin",
			mock.Anything,
			mock.AnythingOfType("admins.Domain")).Return(admins.Domain{}, errors.New("Unexpected Error")).Once()

		admin, err := adminService.UpdateAdmin(context.Background(), admins.Domain{
			Email:    "daaffa@net.usc",
			Password: "kecoak11",
			Name:     "Pablo",
			Address:  "Cambridge",
		})

		assert.Error(t, err)
		assert.Equal(t, admin, admins.Domain{})

		// adminRepository.AssertExpectations(t)
	})
}
