package routes

import (
	"MyProject/ProjectALTA/JobDir/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("api/v1/")
	ev1.GET("users", controllers.GetUserController)
	ev1.POST("users/login", controllers.LoginController)
	ev1.POST("users/register", controllers.RegisterController)
	ev1.GET("users/:userId", controllers.DetailUserController)
	ev1.DELETE("users/:userId", controllers.DeleteUserController)
	ev1.PUT("users/:userId", controllers.UpdateUserController)

	return e
}
