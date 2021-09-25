package routes

import (
	"github.com/labstack/echo/v4"
	"main.go/controllers/categories"
	"main.go/controllers/jobs"
	"main.go/controllers/users"
)

type ControllerList struct {
	UserController     users.UserController
	JobController      jobs.JobController
	CategoryController categories.CategoryController
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
	e.POST("jobs", cl.JobController.CreateJob)
	e.DELETE("jobs", cl.JobController.DeleteAllJobs)
	e.GET("jobs", cl.JobController.GetAllJobs)
	e.GET("jobs/:jobId", cl.JobController.GetJobById)
	e.DELETE("jobs/:jobId", cl.JobController.DeleteJobById)
	e.GET("jobs/result", cl.JobController.SearchJobs)

	//CATEGORY
	e.POST("jobs/categories", cl.CategoryController.CreateCategory)
	e.GET("jobs/categories", cl.CategoryController.GetAllCategory)
	e.GET("jobs/categories/:categoryId", cl.CategoryController.GetCategoryById)
	e.DELETE("jobs/categories/:categoryId", cl.CategoryController.DeleteCategoryById)
}
