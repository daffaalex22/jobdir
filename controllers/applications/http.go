package applications

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"main.go/business/applications"
	"main.go/controllers"
	"main.go/controllers/applications/requests"
	"main.go/controllers/applications/responses"
)

type ApplicationController struct {
	ApplicationUseCase applications.Usecase
}

func NewApplicationController(ApplicationUseCase applications.Usecase) *ApplicationController {
	return &ApplicationController{
		ApplicationUseCase: ApplicationUseCase,
	}
}

// func (ApplicationController ApplicationController) Fillapplications(c echo.Context, category []applications.CategoryDomain) error {

// 	ctx := c.Request().Context()
// 	res, err := ApplicationController.ApplicationUseCase.Fillapplications(ctx, category)

// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.ToListCategoryDomain(res))
// }

func (ApplicationController ApplicationController) CreateApplication(c echo.Context) error {
	fmt.Println("CreateApplication")
	applicationCreate := requests.ApplicationCreate{}
	c.Bind(&applicationCreate)

	ctx := c.Request().Context()
	application, err := ApplicationController.ApplicationUseCase.CreateApplication(ctx, applicationCreate.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(application))
}

func (ApplicationController ApplicationController) GetAllApplications(c echo.Context) error {
	fmt.Println("GetAllapplications")

	ctx := c.Request().Context()
	Application, err := ApplicationController.ApplicationUseCase.GetAllApplications(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ListFromDomain(Application))
}

// func (ApplicationController ApplicationController) GetApplicationById(c echo.Context) error {
// 	fmt.Println("GetApplicationById")

// 	ApplicationId, err := strconv.Atoi(c.Param("ApplicationId"))
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	ctx := c.Request().Context()
// 	Application, err := ApplicationController.ApplicationUseCase.GetApplicationById(ctx, ApplicationId)
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.FromDomain(Application))
// }

func (ApplicationController ApplicationController) DeleteAllApplications(c echo.Context) error {
	fmt.Println("HardDeleteAllApplicationRecords")

	ctx := c.Request().Context()
	error := ApplicationController.ApplicationUseCase.DeleteAllApplications(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(applications.Domain{}))
}

// func (ApplicationController ApplicationController) DeleteApplicationById(c echo.Context) error {
// 	fmt.Println("DeleteApplicationById")

// 	ApplicationId, err := strconv.Atoi(c.Param("ApplicationId"))
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	ctx := c.Request().Context()
// 	Application, err := ApplicationController.ApplicationUseCase.DeleteApplicationById(ctx, ApplicationId)
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.FromDomain(Application))
// }

// func (ApplicationController ApplicationController) Searchapplications(c echo.Context) error {
// 	fmt.Println("GetAllapplications")

// 	title := c.QueryParam("title")

// 	ctx := c.Request().Context()
// 	Application, err := ApplicationController.ApplicationUseCase.Searchapplications(ctx, title)
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.ListFromDomain(Application))
// }

// func (ApplicationController ApplicationController) FilterApplicationByCategory(c echo.Context) error {
// 	fmt.Println("GetAllapplications")

// 	categoryId, err := strconv.Atoi(c.QueryParam("categoryId"))
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	ctx := c.Request().Context()
// 	Application, err := ApplicationController.ApplicationUseCase.FilterApplicationByCategory(ctx, categoryId)
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.ListFromDomain(Application))
// }
