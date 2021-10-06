package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"main.go/business/users"
	"main.go/controllers"
	"main.go/controllers/users/requests"
	"main.go/controllers/users/responses"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(UserUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: UserUseCase,
	}
}

func (UserController UserController) Login(c echo.Context) error {
	fmt.Println("Login")
	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	user, error := UserController.UserUseCase.Login(ctx, userLogin.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomainLogin(user))
}

func (UserController UserController) GetUserById(c echo.Context) error {
	fmt.Println("GetById")

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	user, err := UserController.UserUseCase.GetUserById(ctx, userId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(user))
}

func (UserController UserController) GetAllUser(c echo.Context) error {
	fmt.Println("GetAllUser")

	ctx := c.Request().Context()
	user, err := UserController.UserUseCase.GetAllUser(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ListFromDomain(user))
}

func (UserController UserController) UpdateUser(c echo.Context) error {
	fmt.Println("UpdateUser")

	userUpdate := requests.UserUpdate{}
	c.Bind(&userUpdate)

	ctx := c.Request().Context()
	user, err := UserController.UserUseCase.UpdateUser(ctx, userUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(user))
}

func (UserController UserController) DeleteUser(c echo.Context) error {
	fmt.Println("DeleteUser")

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	user, err := UserController.UserUseCase.DeleteUser(ctx, userId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(user))
}

func (UserController UserController) RegisterUser(c echo.Context) error {
	fmt.Println("RegisterUser")
	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)

	ctx := c.Request().Context()
	user, error := UserController.UserUseCase.RegisterUser(ctx, userRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(user))
}
