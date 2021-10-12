package admins

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"main.go/business/admins"
	"main.go/controllers"
	"main.go/controllers/admins/requests"
	"main.go/controllers/admins/responses"
)

type AdminController struct {
	AdminUseCase admins.Usecase
}

func NewAdminController(AdminUseCase admins.Usecase) *AdminController {
	return &AdminController{
		AdminUseCase: AdminUseCase,
	}
}

func (AdminController AdminController) Login(c echo.Context) error {
	fmt.Println("Login")
	adminLogin := requests.AdminLogin{}
	c.Bind(&adminLogin)

	ctx := c.Request().Context()
	Admin, error := AdminController.AdminUseCase.Login(ctx, adminLogin.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomainLogin(Admin))
}

func (AdminController AdminController) GetAdminById(c echo.Context) error {
	fmt.Println("GetById")

	adminId, err := strconv.Atoi(c.Param("AdminId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	admin, err := AdminController.AdminUseCase.GetAdminById(ctx, adminId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(admin))
}

func (AdminController AdminController) GetAllAdmin(c echo.Context) error {
	fmt.Println("GetAllAdmin")

	ctx := c.Request().Context()
	admin, err := AdminController.AdminUseCase.GetAllAdmin(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ListFromDomain(admin))
}

func (AdminController AdminController) UpdateAdmin(c echo.Context) error {
	fmt.Println("UpdateAdmin")

	adminUpdate := requests.AdminUpdate{}
	c.Bind(&adminUpdate)

	ctx := c.Request().Context()
	admin, err := AdminController.AdminUseCase.UpdateAdmin(ctx, adminUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(admin))
}

func (AdminController AdminController) DeleteAdmin(c echo.Context) error {
	fmt.Println("DeleteAdmin")

	adminId, err := strconv.Atoi(c.Param("adminId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	admin, err := AdminController.AdminUseCase.DeleteAdmin(ctx, adminId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(admin))
}

func (AdminController AdminController) RegisterAdmin(c echo.Context) error {
	fmt.Println("RegisterAdmin")
	adminRegister := requests.AdminRegister{}
	c.Bind(&adminRegister)

	ctx := c.Request().Context()
	admin, error := AdminController.AdminUseCase.RegisterAdmin(ctx, adminRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(admin))
}
