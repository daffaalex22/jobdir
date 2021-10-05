package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"main.go/app/middlewares"
	"main.go/app/routes"
	_adminUsecase "main.go/business/admins"
	_applicationUsecase "main.go/business/applications"
	_categoryUsecase "main.go/business/categories"
	_companyUsecase "main.go/business/companies"
	_jobUsecase "main.go/business/jobs"
	_userUsecase "main.go/business/users"
	_adminController "main.go/controllers/admins"
	_applicationController "main.go/controllers/applications"
	_categoryController "main.go/controllers/categories"
	_companyController "main.go/controllers/companies"
	_jobController "main.go/controllers/jobs"
	_userController "main.go/controllers/users"
	_admindb "main.go/drivers/databases/admins"
	_applicationdb "main.go/drivers/databases/applications"
	_categorydb "main.go/drivers/databases/categories"
	_companydb "main.go/drivers/databases/companies"
	_jobdb "main.go/drivers/databases/jobs"
	_userdb "main.go/drivers/databases/users"
	_mysqlDriver "main.go/drivers/mysql"
	// "main.go/helpers/mongowriter"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userdb.Users{})
	db.AutoMigrate(&_jobdb.Jobs{})
	db.AutoMigrate(&_categorydb.Categories{})
	db.AutoMigrate(&_admindb.Admins{})
	db.AutoMigrate(&_companydb.Companies{})
	db.AutoMigrate(&_applicationdb.Applications{})
}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	// configMongo := _mongoDriver.ConfigMongo{
	// 	DB_Host: viper.GetString(`logger.host`),
	// }

	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	// configLogger := middlewares.ConfigLogger{
	// 	Format: viper.GetString(`logger.format`),
	// }

	// Session := configMongo.InitMongo()
	// mw := &mongowriter.MongoWriter{
	// 	Session: Session,
	// }

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	// configLogger.MongoOut(e, mw)

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}",
	// }))
	// e.Logger.SetOutput(mw)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userdb.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext, &configJWT)
	userController := _userController.NewUserController(userUseCase)

	jobRepository := _jobdb.NewMysqlJobRepository(Conn)
	jobUseCase := _jobUsecase.NewJobUsecase(jobRepository, timeoutContext)
	jobController := _jobController.NewJobController(jobUseCase)

	categoryRepository := _categorydb.NewMysqlCategoryRepository(Conn)
	categoryUseCase := _categoryUsecase.NewCategoryUsecase(categoryRepository, timeoutContext)
	categoryController := _categoryController.NewCategoryController(categoryUseCase)

	adminRepository := _admindb.NewMysqlAdminRepository(Conn)
	adminUseCase := _adminUsecase.NewAdminUsecase(adminRepository, timeoutContext, &configJWT)
	adminController := _adminController.NewAdminController(adminUseCase)

	companyRepository := _companydb.NewMysqlCompanyRepository(Conn)
	companyUseCase := _companyUsecase.NewCompanyUsecase(companyRepository, timeoutContext)
	companyController := _companyController.NewCompanyController(companyUseCase)

	applicationRepository := _applicationdb.NewMysqlApplicationRepository(Conn)
	applicationUseCase := _applicationUsecase.NewApplicationUsecase(applicationRepository, timeoutContext)
	applicationController := _applicationController.NewApplicationController(applicationUseCase)

	routesInit := routes.ControllerList{
		// LoggerConfig:          configLogger.Init(),
		JwtConfig:             configJWT.Init(),
		UserController:        *userController,
		JobController:         *jobController,
		CategoryController:    *categoryController,
		AdminController:       *adminController,
		CompanyController:     *companyController,
		ApplicationController: *applicationController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
