package routes

import (
	"github.com/labstack/echo/v4"
	"main.go/controllers/jobs"
	"main.go/controllers/users"
)

type ControllerList struct {
	UserController users.UserController
	JobController  jobs.JobController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	// USER
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.RegisterUser)
	e.GET("users/:userId", cl.UserController.GetUserById)
	e.GET("users", cl.UserController.GetAllUser)
	e.PUT("users", cl.UserController.UpdateUser)
	e.DELETE("users/:userId", cl.UserController.DeleteUser)

	// JOB
	e.POST("jobs/post", cl.JobController.CreateJob)
	e.DELETE("jobs", cl.JobController.DeleteAllJobs)
	e.GET("jobs", cl.JobController.GetAllJobs)
	e.GET("jobs/:jobId", cl.JobController.GetJobById)
	e.DELETE("jobs/:jobId", cl.JobController.DeleteJobById)
}
