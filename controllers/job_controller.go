package controllers

import (
	"MyProject/ProjectALTA/JobDir/configs"
	"MyProject/ProjectALTA/JobDir/models/jobs"
	"MyProject/ProjectALTA/JobDir/models/response"
	"MyProject/ProjectALTA/JobDir/models/users"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetJobController(c echo.Context) error {

	jobs := []jobs.Job{}

	result := configs.DB.Find(&jobs)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data jobs dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data jobs",
		Data:    jobs,
	})
}

func DetailJobController(c echo.Context) error {
	jobId, err := strconv.Atoi(c.Param("jobId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi jobId",
			Data:    nil,
		})
	}

	job := users.User{}

	result := configs.DB.First(&job, jobId)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data job dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan detail data job",
		Data:    job,
	})
}
