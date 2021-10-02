package users_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/business/users"
	_mockUserRepository "main.go/business/users/mocks"
)

var userRepository _mockUserRepository.Repository
var userservice users.Usecase
var userDomain users.Domain
var usersDomain []users.Domain

// var configJWT *middlewares.ConfigJWT

func setup() {
	userservice = users.NewUserUsecase(&userRepository, time.Hour*1 /*, configJWT */)
	userDomain = users.Domain{
		Id:      1,
		Name:    "Pabby",
		Email:   "daaffa@net.usc",
		Address: "Belanda",
		// JobsCreated: []jobs.Domain
		Password: "kecoak11",
		Token:    "123",
	}
	usersDomain = append(usersDomain, userDomain)
}

func TestLogin(t *testing.T) {

	setup()
	userRepository.On("Login",
		mock.Anything,
		mock.AnythingOfType("users.Domain")).Return(userDomain, nil).Once()

	t.Run("Test case 1 | Valid Login", func(t *testing.T) {
		user, err := userservice.Login(context.Background(), users.Domain{
			Email:    "daaffa@net.usc",
			Password: "kecoak11",
		})

		assert.Nil(t, err)
		assert.Equal(t, "Pabby", user.Name)
	})

	t.Run("Test Case 2 | Invalid Email Empty", func(t *testing.T) {

		_, err := userservice.Login(context.Background(), users.Domain{
			Email:    "",
			Password: "kecoak11",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Password Empty", func(t *testing.T) {

		_, err := userservice.Login(context.Background(), users.Domain{
			Email:    "daaffa@net.usc",
			Password: "",
		})
		assert.NotNil(t, err)
	})

}

func TestGetuserById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("GetUserById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		user, err := userservice.GetUserById(context.Background(), userDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		userRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	userRepository.On("GetuserById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(userDomain, nil).Once()

	// 	user, err := userservice.GetuserById(context.Background(), userDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, users.Domain{}, user)

	// 	userRepository.AssertExpectations(t)
	// })
}

func TestGetAlluser(t *testing.T) {
	setup()
	// usersDomain = append(usersDomain, userDomain)

	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("GetAllUser",
			mock.Anything).Return(usersDomain, nil).Once()

		user, err := userservice.GetAllUser(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)

		userRepository.AssertExpectations(t)
	})
}

func TestDeleteuser(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("DeleteUser",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		user, err := userservice.DeleteUser(context.Background(), userDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		userRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	userRepository.On("GetuserById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(userDomain, nil).Once()

	// 	user, err := userservice.GetuserById(context.Background(), userDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, users.Domain{}, user)

	// 	userRepository.AssertExpectations(t)
	// })
}

func TestRegisterUser(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Registry", func(t *testing.T) {
		userRepository.On("RegisterUser",
			mock.Anything,
			mock.AnythingOfType("users.Domain")).Return(userDomain, nil).Once()
		user, err := userservice.RegisterUser(context.Background(), users.Domain{
			Email:    "daaffa@net.usc",
			Password: "kecoak11",
		})

		assert.Nil(t, err)
		assert.Equal(t, "Pabby", user.Name)
	})

	t.Run("Test Case 2 | Invalid Email Empty", func(t *testing.T) {
		userRepository.On("RegisterUser",
			mock.Anything,
			mock.AnythingOfType("users.Domain")).Return(userDomain, nil).Once()
		_, err := userservice.RegisterUser(context.Background(), users.Domain{
			Email:    "",
			Password: "kecoak11",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Password Empty", func(t *testing.T) {
		userRepository.On("RegisterUser",
			mock.Anything,
			mock.AnythingOfType("users.Domain")).Return(userDomain, nil).Once()
		_, err := userservice.RegisterUser(context.Background(), users.Domain{
			Email:    "daaffa@net.usc",
			Password: "",
		})
		assert.NotNil(t, err)
	})

}
