package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models/jobs"
	"main.go/models/users"
)

var DB *gorm.DB

// type Configration struct {
// 	DB_USERNAME string
// 	DB_PASSWORD string
// 	DB_PORT     string
// 	DB_HOST     string
// 	DB_NAME     string
// }

func InitDB() {
	dsn := "root:haji12345@tcp(127.0.0.1:3306)/jobdir?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&users.User{}, &jobs.Job{}, &jobs.JobDesc{})
}
