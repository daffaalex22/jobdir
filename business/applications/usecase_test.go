package applications_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/business/applications"
	_mockApplicationRepository "main.go/business/applications/mocks"
)

var ApplicationRepository _mockApplicationRepository.Repository
var applicationservice applications.Usecase
var ApplicationDomain applications.Domain
var applicationsDomain []applications.Domain

// var configJWT *middlewares.ConfigJWT

func setup() {
	applicationservice = applications.NewApplicationUsecase(&ApplicationRepository, time.Hour*1 /*, configJWT */)
	ApplicationDomain = applications.Domain{
		UserId: 1,
		JobId:  1,
		Status: "APPLYING",
	}
	applicationsDomain = append(applicationsDomain, ApplicationDomain)
}

func TestGetAllApplications(t *testing.T) {
	setup()
	// applicationsDomain = append(applicationsDomain, ApplicationDomain)

	t.Run("Test case 1 | Valid GetAll", func(t *testing.T) {
		ApplicationRepository.On("GetAllApplications",
			mock.Anything).Return(applicationsDomain, nil).Once()

		Application, err := applicationservice.GetAllApplications(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, Application)

		ApplicationRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error GetAll", func(t *testing.T) {
		ApplicationRepository.On("GetAllApplications",
			mock.Anything).Return([]applications.Domain{}, errors.New(mock.Anything)).Once()

		Application, err := applicationservice.GetAllApplications(context.Background())

		assert.Error(t, err)
		assert.Equal(t, Application, []applications.Domain{})

		ApplicationRepository.AssertExpectations(t)
	})
}

// func TestDeleteAllApplications(t *testing.T) {
// 	setup()

// 	t.Run("Test case 1", func(t *testing.T) {
// 		ApplicationRepository.On("DeleteAllApplications",
// 			mock.Anything,
// 			mock.AnythingOfType("int")).Return(ApplicationDomain, nil).Once()

// 		err := applicationservice.DeleteAllApplications(context.Background())

// 		assert.NoError(t, err)
// 		// assert.NotNil(t, Application)

// 		ApplicationRepository.AssertExpectations(t)
// 	})

// 	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
// 	// 	ApplicationRepository.On("GetApplicationById",
// 	// 		mock.Anything,
// 	// 		mock.AnythingOfType("int")).Return(ApplicationDomain, nil).Once()

// 	// 	Application, err := applicationservice.GetApplicationById(context.Background(), ApplicationDomain.Id)

// 	// 	assert.Error(t, err)
// 	// 	assert.Equal(t, applications.Domain{}, Application)

// 	// 	ApplicationRepository.AssertExpectations(t)
// 	// })
// }

func TestCreateApplication(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Registry", func(t *testing.T) {
		ApplicationRepository.On("CreateApplication",
			mock.Anything,
			mock.AnythingOfType("applications.Domain")).Return(ApplicationDomain, nil).Once()
		application, err := applicationservice.CreateApplication(context.Background(), applications.Domain{
			UserId: 1,
			JobId:  1,
			Status: "APPLYING",
		})

		assert.Nil(t, err)
		assert.NotNil(t, application)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		ApplicationRepository.On("CreateApplication",
			mock.Anything,
			mock.AnythingOfType("applications.Domain")).Return(applications.Domain{}, errors.New(mock.Anything)).Once()
		application, err := applicationservice.CreateApplication(context.Background(), applications.Domain{
			UserId: 1,
			JobId:  1,
			Status: "APPLYING",
		})

		assert.NotNil(t, err)
		assert.Equal(t, application, applications.Domain{})
	})

	// t.Run("Test Case 2 | Invalid Email Empty", func(t *testing.T) {
	// 	ApplicationRepository.On("CreateApplication",
	// 		mock.Anything,
	// 		mock.AnythingOfType("applications.Domain")).Return(ApplicationDomain, nil).Once()
	// 	_, err := applicationservice.CreateApplication(context.Background(), applications.Domain{
	// 		Email:    "",
	// 		Password: "kecoak11",
	// 	})
	// 	assert.NotNil(t, err)
	// })

	// t.Run("Test Case 3 | Invalid Password Empty", func(t *testing.T) {
	// 	ApplicationRepository.On("CreateApplication",
	// 		mock.Anything,
	// 		mock.AnythingOfType("applications.Domain")).Return(ApplicationDomain, nil).Once()
	// 	_, err := applicationservice.CreateApplication(context.Background(), applications.Domain{
	// 		Email:    "daaffa@net.usc",
	// 		Password: "",
	// 	})
	// 	assert.NotNil(t, err)
	// })

}

func TestDeleteAllApplication(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		ApplicationRepository.On("DeleteAllApplications",
			mock.Anything).Return(nil).Once()

		err := applicationservice.DeleteAllApplications(context.Background())

		assert.NoError(t, err)

		// ApplicationRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		ApplicationRepository.On("DeleteAllApplications",
			mock.Anything).Return(errors.New(mock.Anything)).Once()

		err := applicationservice.DeleteAllApplications(context.Background())

		assert.NotNil(t, err)

		// ApplicationRepository.AssertExpectations(t)
	})
}
