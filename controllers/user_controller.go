package controllers

import (
	"MyProject/ProjectALTA/JobDir/configs"
	"MyProject/ProjectALTA/JobDir/models/response"
	"MyProject/ProjectALTA/JobDir/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"main.go/middlewares"
)

func RegisterController(c echo.Context) error {
	var userRegister users.UserRegister
	c.Bind(&userRegister)

	// validasi
	if userRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Nama masih kosong",
			Data:    nil,
		})
	}

	// etc

	var userDB users.User
	userDB.Name = userRegister.Name
	userDB.Password = userRegister.Password
	userDB.Address = userRegister.Address
	userDB.Email = userRegister.Email

	result := configs.DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data user ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil register",
		Data:    userDB,
	})
}

func UpdateUserController(c echo.Context) error {
	var userUpdate users.UserRegister
	c.Bind(&userUpdate)

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi userId",
			Data:    nil,
		})
	}

	userDB := users.User{}
	result := configs.DB.First(&userDB, userId)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika mencari data di DB",
			Data:    nil,
		})
	}

	configs.DB.Model(&userDB).Updates(users.User{
		Name:     userUpdate.Name,
		Password: userUpdate.Password,
		Address:  userUpdate.Address,
		Email:    userUpdate.Email})

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil update data user",
		Data:    userDB,
	})
}

func LoginController(c echo.Context) error {
	userLogin := users.UserLogin{}
	c.Bind(&userLogin)
	// login
	user := users.User{}
	result := configs.DB.Where("email = ? AND password = ?", userLogin.Email, userLogin.Password).First(&user)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika mengecek database untuk body yang dikirim",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusForbidden, response.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "Username atau Password Salah!",
			Data:    nil,
		})
	}

	token, err := middlewares.GenerateTokenJWT(user.Id)

	if err != nil {
		return c.JSON(http.StatusForbidden, response.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "Error ketika membuat Token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Login Berhasil! Token: " + token,
		Data:    userLogin,
	})
}

func GetUserController(c echo.Context) error {

	users := []users.User{}

	result := configs.DB.Find(&users)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data user",
		Data:    users,
	})
}

func DetailUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi userId",
			Data:    nil,
		})
	}

	users := users.User{}

	result := configs.DB.First(&users, userId)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan detail data user",
		Data:    users,
	})
}

func DeleteUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi userId",
			Data:    nil,
		})
	}

	users := users.User{}
	result := configs.DB.First(&users, userId)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data tersebut di DB",
			Data:    nil,
		})
	}

	result = configs.DB.Delete(&users, []int{userId})

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil menghapus data user berikut",
		Data:    users,
	})
}
