package categories

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"main.go/business/categories"
	"main.go/controllers"
	"main.go/controllers/categories/requests"
	"main.go/controllers/categories/responses"
)

type CategoryController struct {
	CategoryUseCase categories.Usecase
}

func NewCategoryController(CategoryUseCase categories.Usecase) *CategoryController {
	return &CategoryController{
		CategoryUseCase: CategoryUseCase,
	}
}

func (CategoryController CategoryController) CreateCategory(c echo.Context) error {
	fmt.Println("CreateCategory")
	categoryCreate := requests.CategoryCreate{}
	c.Bind(&categoryCreate)

	ctx := c.Request().Context()
	category, err := CategoryController.CategoryUseCase.CreateCategory(ctx, categoryCreate.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(category))
}

func (CategoryController CategoryController) GetAllCategory(c echo.Context) error {
	fmt.Println("GetAllCategories")

	ctx := c.Request().Context()
	category, err := CategoryController.CategoryUseCase.GetAllCategory(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ListFromDomain(category))
}

func (CategoryController CategoryController) GetCategoryById(c echo.Context) error {
	fmt.Println("GetCategoryById")

	categoryId, err := strconv.Atoi(c.Param("CategoryId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	category, err := CategoryController.CategoryUseCase.GetCategoryById(ctx, categoryId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(category))
}

// func (CategoryController CategoryController) DeleteAllcategories(c echo.Context) error {
// 	fmt.Println("HardDeleteAllCategoryRecords")

// 	ctx := c.Request().Context()
// 	error := CategoryController.CategoryUseCase.DeleteAllcategories(ctx)

// 	if error != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.FromDomain(categories.Domain{}))
// }

func (CategoryController CategoryController) DeleteCategoryById(c echo.Context) error {
	fmt.Println("DeleteCategoryById")

	categoryId, err := strconv.Atoi(c.Param("CategoryId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	category, err := CategoryController.CategoryUseCase.DeleteCategoryById(ctx, categoryId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(category))
}
