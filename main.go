package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"main.go/app/routes"
	_categoryUsecase "main.go/business/categories"
	_jobUsecase "main.go/business/jobs"
	_userUsecase "main.go/business/users"
	_categoryController "main.go/controllers/categories"
	_jobController "main.go/controllers/jobs"
	_userController "main.go/controllers/users"
	_categorydb "main.go/drivers/databases/categories"
	_jobdb "main.go/drivers/databases/jobs"
	_userdb "main.go/drivers/databases/users"
	_mysqlDriver "main.go/drivers/mysql"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
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
}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userdb.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	jobRepository := _jobdb.NewMysqlJobRepository(Conn)
	jobUseCase := _jobUsecase.NewJobUsecase(jobRepository, timeoutContext)
	jobController := _jobController.NewJobController(jobUseCase)

	categoryRepository := _categorydb.NewMysqlCategoryRepository(Conn)
	categoryUseCase := _categoryUsecase.NewCategoryUsecase(categoryRepository, timeoutContext)
	categoryController := _categoryController.NewCategoryController(categoryUseCase)

	routesInit := routes.ControllerList{
		UserController:     *userController,
		JobController:      *jobController,
		CategoryController: *categoryController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
