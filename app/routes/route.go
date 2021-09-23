package routes

import (
	"github.com/labstack/echo/v4"
	"main.go/controllers/users"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/login", cl.UserController.Login)
	e.GET("users/:userId", cl.UserController.GetUserById)
	e.GET("users", cl.UserController.GetAllUser)
	e.PUT("users", cl.UserController.UpdateUser)
}
