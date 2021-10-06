package jobs_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/business/jobs"
	_mockJobRepository "main.go/business/jobs/mocks"
)

var JobRepository _mockJobRepository.Repository
var jobService jobs.Usecase
var jobDomain jobs.Domain
var jobsDomain []jobs.Domain

// var configJWT *middlewares.ConfigJWT

func setup() {
	jobService = jobs.NewJobUsecase(&JobRepository, time.Hour*1 /*, configJWT */)
	jobDomain = jobs.Domain{
		Id:         1,
		Title:      "Software Engineer",
		CategoryId: 1,
		JobDesc:    "Kerja Lembur Bagai Kuda",
		CreatedBy:  1,
		CompanyId:  1,
	}
	jobsDomain = append(jobsDomain, jobDomain)
}

func TestGetJobById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		JobRepository.On("GetJobById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(jobDomain, nil).Once()

		Job, err := jobService.GetJobById(context.Background(), jobDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, Job)

		// JobRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		JobRepository.On("GetJobById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(jobs.Domain{}, errors.New(mock.Anything)).Once()

		Job, err := jobService.GetJobById(context.Background(), jobDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, Job, jobs.Domain{})

		// JobRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	JobRepository.On("GetJobById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(jobDomain, nil).Once()

	// 	Job, err := jobService.GetJobById(context.Background(), jobDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, jobs.Domain{}, Job)

	JobRepository.AssertExpectations(t)
	// })
}

func TestGetAllJobs(t *testing.T) {
	setup()
	// jobsDomain = append(jobsDomain, jobDomain)

	t.Run("Test case 1", func(t *testing.T) {
		JobRepository.On("GetAllJobs",
			mock.Anything).Return(jobsDomain, nil).Once()

		Job, err := jobService.GetAllJobs(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, Job)

		// JobRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		JobRepository.On("GetAllJobs",
			mock.Anything).Return([]jobs.Domain{}, errors.New(mock.Anything)).Once()

		Job, err := jobService.GetAllJobs(context.Background())

		assert.Error(t, err)
		assert.Equal(t, Job, []jobs.Domain{})

		// JobRepository.AssertExpectations(t)
	})
}

func TestDeleteJobById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		JobRepository.On("DeleteJobById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(jobDomain, nil).Once()

		Job, err := jobService.DeleteJobById(context.Background(), jobDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, Job)

		// JobRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Return Error", func(t *testing.T) {
		JobRepository.On("DeleteJobById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(jobs.Domain{}, errors.New(mock.Anything)).Once()

		Job, err := jobService.DeleteJobById(context.Background(), jobDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, Job, jobs.Domain{})

		// JobRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	JobRepository.On("GetJobById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(jobDomain, nil).Once()

	// 	Job, err := jobService.GetJobById(context.Background(), jobDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, jobs.Domain{}, Job)

	// JobRepository.AssertExpectations(t)
	// })
}

func TestCreateJob(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Create", func(t *testing.T) {
		JobRepository.On("CreateJob",
			mock.Anything,
			mock.AnythingOfType("jobs.Domain")).Return(jobDomain, nil).Once()
		Job, err := jobService.CreateJob(context.Background(), jobs.Domain{
			Id:         1,
			Title:      "Software Engineer",
			CategoryId: 1,
			JobDesc:    "Kerja Lembur Bagai Kuda",
			CreatedBy:  1,
			CompanyId:  1,
		})

		assert.Nil(t, err)
		assert.NotNil(t, Job)
	})

	t.Run("Test Case 2 | Return Error", func(t *testing.T) {
		JobRepository.On("CreateJob",
			mock.Anything,
			mock.AnythingOfType("jobs.Domain")).Return(jobs.Domain{}, errors.New(mock.Anything)).Once()
		Job, err := jobService.CreateJob(context.Background(), jobs.Domain{
			Id:         1,
			Title:      "Software Engineer",
			CategoryId: 1,
			JobDesc:    "Kerja Lembur Bagai Kuda",
			CreatedBy:  1,
			CompanyId:  1,
		})
		assert.Error(t, err)
		assert.Equal(t, Job, jobs.Domain{})
	})

	t.Run("Test Case 3 | Invalid Job Empty", func(t *testing.T) {
		JobRepository.On("CreateJob",
			mock.Anything,
			mock.AnythingOfType("jobs.Domain")).Return(jobDomain, nil).Once()
		_, err := jobService.CreateJob(context.Background(), jobs.Domain{
			Id:         1,
			Title:      "",
			CategoryId: 1,
			JobDesc:    "Kerja Lembur Bagai Kuda",
			CreatedBy:  1,
			CompanyId:  1,
		})
		assert.NotNil(t, err)
	})

}

func TestDeleteAllJobs(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		JobRepository.On("DeleteAllJobs",
			mock.Anything).Return(nil).Once()

		err := jobService.DeleteAllJobs(context.Background())

		assert.NoError(t, err)

		// JobRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		JobRepository.On("DeleteAllJobs",
			mock.Anything).Return(errors.New(mock.Anything)).Once()

		err := jobService.DeleteAllJobs(context.Background())

		assert.Error(t, err)

		// JobRepository.AssertExpectations(t)
	})
}
func TestSearchJobs(t *testing.T) {
	setup()
	// jobsDomain = append(jobsDomain, jobDomain)

	t.Run("Test case 1", func(t *testing.T) {
		JobRepository.On("SearchJobs",
			mock.Anything,
			mock.AnythingOfType("string")).Return(jobsDomain, nil).Once()

		Job, err := jobService.SearchJobs(context.Background(), "soft")

		assert.NoError(t, err)
		assert.NotNil(t, Job)

		// JobRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		JobRepository.On("SearchJobs",
			mock.Anything,
			mock.AnythingOfType("string")).Return([]jobs.Domain{}, errors.New(mock.Anything)).Once()

		Job, err := jobService.SearchJobs(context.Background(), "soft")

		assert.Error(t, err)
		assert.Equal(t, Job, []jobs.Domain{})

		// JobRepository.AssertExpectations(t)
	})
}

func TestFilterJobByCategory(t *testing.T) {
	setup()
	// jobsDomain = append(jobsDomain, jobDomain)

	t.Run("Test case 1", func(t *testing.T) {
		JobRepository.On("FilterJobByCategory",
			mock.Anything,
			mock.AnythingOfType("int")).Return(jobsDomain, nil).Once()

		Job, err := jobService.FilterJobByCategory(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, Job)

		// JobRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		JobRepository.On("FilterJobByCategory",
			mock.Anything,
			mock.AnythingOfType("int")).Return([]jobs.Domain{}, errors.New(mock.Anything)).Once()

		Job, err := jobService.FilterJobByCategory(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, Job, []jobs.Domain{})

		// JobRepository.AssertExpectations(t)
	})
}
