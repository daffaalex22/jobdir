package users

import (
	"fmt"
	"net/http"

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
	user, error := UserController.UserUseCase.Login(ctx, userLogin.Email, userLogin.Password)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(user))
}
