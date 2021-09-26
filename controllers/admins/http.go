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
	AdminLogin := requests.AdminLogin{}
	c.Bind(&AdminLogin)

	ctx := c.Request().Context()
	Admin, error := AdminController.AdminUseCase.Login(ctx, AdminLogin.Email, AdminLogin.Password)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(Admin))
}

func (AdminController AdminController) GetAdminById(c echo.Context) error {
	fmt.Println("GetById")

	AdminId, err := strconv.Atoi(c.Param("AdminId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	Admin, err := AdminController.AdminUseCase.GetAdminById(ctx, AdminId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(Admin))
}

func (AdminController AdminController) GetAllAdmin(c echo.Context) error {
	fmt.Println("GetAllAdmin")

	ctx := c.Request().Context()
	Admin, err := AdminController.AdminUseCase.GetAllAdmin(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ListFromDomain(Admin))
}

func (AdminController AdminController) UpdateAdmin(c echo.Context) error {
	fmt.Println("UpdateAdmin")

	AdminUpdate := requests.AdminUpdate{}
	c.Bind(&AdminUpdate)

	ctx := c.Request().Context()
	Admin, err := AdminController.AdminUseCase.UpdateAdmin(ctx, AdminUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(Admin))
}

func (AdminController AdminController) DeleteAdmin(c echo.Context) error {
	fmt.Println("DeleteAdmin")

	AdminId, err := strconv.Atoi(c.Param("AdminId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	Admin, err := AdminController.AdminUseCase.DeleteAdmin(ctx, AdminId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(Admin))
}

func (AdminController AdminController) RegisterAdmin(c echo.Context) error {
	fmt.Println("RegisterAdmin")
	AdminRegister := requests.AdminRegister{}
	c.Bind(&AdminRegister)

	ctx := c.Request().Context()
	Admin, error := AdminController.AdminUseCase.RegisterAdmin(ctx, AdminRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(Admin))
}

func (AdminController AdminController) HardDeleteAllAdmins(c echo.Context) error {
	fmt.Println("HardDeleteAllAdmins")

	ctx := c.Request().Context()
	err := AdminController.AdminUseCase.HardDeleteAllAdmins(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(admins.Domain{}))
}
