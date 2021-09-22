package routes

import (
	"MyProject/ProjectALTA/JobDir/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main.go/constants"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	jwt := middleware.JWT([]byte(constants.SECRET_JWT_USER))

	ev1 := e.Group("api/v1/")

	// User
	ev1.GET("users", controllers.GetUserController, jwt)
	ev1.POST("users/login", controllers.LoginController)
	ev1.POST("users/register", controllers.RegisterController)
	ev1.GET("users/:userId", controllers.DetailUserController, jwt)
	ev1.DELETE("users/:userId", controllers.DeleteUserController)
	ev1.PUT("users/:userId", controllers.UpdateUserController)

	// Job
	ev1.GET("jobs", controllers.GetJobController)
	ev1.GET("jobs/:jobId", controllers.DetailJobController)
	ev1.POST("jobs", controllers.PostJobController)
	ev1.DELETE("jobs", controllers.DeleteAllJobsController)

	// Job Desc
	ev1.GET("jobdescs", controllers.GetJobDescController)
	return e
}
