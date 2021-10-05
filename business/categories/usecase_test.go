package categories_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/business/categories"
	_mockCategoryRepository "main.go/business/categories/mocks"
)

var CategoryRepository _mockCategoryRepository.Repository
var categorieService categories.Usecase
var categoryDomain categories.Domain
var categoriesDomain []categories.Domain

// var configJWT *middlewares.ConfigJWT

func setup() {
	categorieService = categories.NewCategoryUsecase(&CategoryRepository, time.Hour*1 /*, configJWT */)
	categoryDomain = categories.Domain{
		Id:       1,
		Category: "Software Development",
	}
	categoriesDomain = append(categoriesDomain, categoryDomain)
}

func TestGetCategoryById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		CategoryRepository.On("GetCategoryById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(categoryDomain, nil).Once()

		Category, err := categorieService.GetCategoryById(context.Background(), categoryDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, Category)

		CategoryRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		CategoryRepository.On("GetCategoryById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(categories.Domain{}, errors.New(mock.Anything)).Once()

		category, err := categorieService.GetCategoryById(context.Background(), categoryDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, category, categories.Domain{})

		CategoryRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	CategoryRepository.On("GetCategoryById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(CategoryDomain, nil).Once()

	// 	Category, err := categorieService.GetCategoryById(context.Background(), CategoryDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, categories.Domain{}, Category)

	// 	CategoryRepository.AssertExpectations(t)
	// })
}

func TestGetAllCategory(t *testing.T) {
	setup()
	// categoriesDomain = append(categoriesDomain, CategoryDomain)

	t.Run("Test case 1", func(t *testing.T) {
		CategoryRepository.On("GetAllCategory",
			mock.Anything).Return(categoriesDomain, nil).Once()

		Category, err := categorieService.GetAllCategory(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, Category)

		CategoryRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		CategoryRepository.On("GetAllCategory",
			mock.Anything).Return([]categories.Domain{}, errors.New(mock.Anything)).Once()

		category, err := categorieService.GetAllCategory(context.Background())

		assert.Error(t, err)
		assert.Equal(t, category, []categories.Domain{})

		CategoryRepository.AssertExpectations(t)
	})
}

func TestDeleteCategoryById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		CategoryRepository.On("DeleteCategoryById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(categoryDomain, nil).Once()

		Category, err := categorieService.DeleteCategoryById(context.Background(), categoryDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, Category)

		CategoryRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Return Error", func(t *testing.T) {
		CategoryRepository.On("DeleteCategoryById",
			mock.Anything,
			mock.AnythingOfType("int")).Return(categories.Domain{}, errors.New(mock.Anything)).Once()

		category, err := categorieService.DeleteCategoryById(context.Background(), categoryDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, category, categories.Domain{})

		CategoryRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | Error Failed", func(t *testing.T) {
	// 	CategoryRepository.On("GetCategoryById",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(CategoryDomain, nil).Once()

	// 	Category, err := categorieService.GetCategoryById(context.Background(), CategoryDomain.Id)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, categories.Domain{}, Category)

	// 	CategoryRepository.AssertExpectations(t)
	// })
}

func TestCreateCategory(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Create", func(t *testing.T) {
		CategoryRepository.On("CreateCategory",
			mock.Anything,
			mock.AnythingOfType("categories.Domain")).Return(categoryDomain, nil).Once()
		category, err := categorieService.CreateCategory(context.Background(), categories.Domain{
			Category: "Software Development",
		})

		assert.Nil(t, err)
		assert.NotNil(t, category)
	})

	t.Run("Test Case 2 | Return Error", func(t *testing.T) {
		CategoryRepository.On("CreateCategory",
			mock.Anything,
			mock.AnythingOfType("categories.Domain")).Return(categories.Domain{}, errors.New(mock.Anything)).Once()
		category, err := categorieService.CreateCategory(context.Background(), categories.Domain{
			Category: "Software Development",
		})
		assert.Error(t, err)
		assert.Equal(t, category, categories.Domain{})
	})

	t.Run("Test Case 3 | Invalid Category Empty", func(t *testing.T) {
		CategoryRepository.On("CreateCategory",
			mock.Anything,
			mock.AnythingOfType("categories.Domain")).Return(categoryDomain, nil).Once()
		_, err := categorieService.CreateCategory(context.Background(), categories.Domain{
			Category: "",
		})
		assert.NotNil(t, err)
	})

}

// func TestHardDeleteAllCategory(t *testing.T) {
// 	setup()

// 	t.Run("Test case 1", func(t *testing.T) {
// 		CategoryRepository.On("HardDeleteAllcategories",
// 			mock.Anything).Return(nil).Once()

// 		err := categorieService.HardDeleteAllcategories(context.Background())

// 		assert.NoError(t, err)

// 		// CategoryRepository.AssertExpectations(t)
// 	})
// }
