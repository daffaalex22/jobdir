package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"main.go/controllers/admins"
	"main.go/controllers/categories"
	"main.go/controllers/jobs"
	"main.go/controllers/users"
)

type ControllerList struct {
	UserController     users.UserController
	AdminController    admins.AdminController
	JobController      jobs.JobController
	CategoryController categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	jwt := middleware.JWT([]byte(viper.GetString(`jwt.secret`)))

	// USER
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.RegisterUser)
	e.GET("users/:userId", cl.UserController.GetUserById, jwt)
	e.GET("users", cl.UserController.GetAllUser, jwt)
	e.PUT("users", cl.UserController.UpdateUser)
	e.DELETE("users/:userId", cl.UserController.DeleteUser)

	// ADMIN
	e.POST("admins/login", cl.AdminController.Login)
	e.POST("admins/register", cl.AdminController.RegisterAdmin)
	e.GET("admins/:userId", cl.AdminController.GetAdminById, jwt)
	e.GET("admins", cl.AdminController.GetAllAdmin, jwt)
	e.PUT("admins", cl.AdminController.UpdateAdmin)
	e.DELETE("admins/:adminId", cl.AdminController.DeleteAdmin)
	e.DELETE("admins", cl.AdminController.HardDeleteAllAdmins)

	// JOB
	e.POST("jobs", cl.JobController.CreateJob)
	e.DELETE("jobs", cl.JobController.DeleteAllJobs)
	e.GET("jobs", cl.JobController.GetAllJobs)
	e.GET("jobs/:jobId", cl.JobController.GetJobById)
	e.DELETE("jobs/:jobId", cl.JobController.DeleteJobById)
	e.GET("jobs/result", cl.JobController.SearchJobs)
	e.GET("jobs/result", cl.JobController.FilterJobByCategory)

	//CATEGORY
	e.POST("jobs/categories", cl.CategoryController.CreateCategory)
	e.GET("jobs/categories", cl.CategoryController.GetAllCategory)
	e.GET("jobs/categories/:categoryId", cl.CategoryController.GetCategoryById)
	e.DELETE("jobs/categories/:categoryId", cl.CategoryController.DeleteCategoryById)
}
