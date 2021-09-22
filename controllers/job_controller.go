package controllers

import (
	"MyProject/ProjectALTA/JobDir/configs"
	"MyProject/ProjectALTA/JobDir/models/jobs"
	"MyProject/ProjectALTA/JobDir/models/response"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PostJobController(c echo.Context) error {
	var jobPost jobs.JobPost
	c.Bind(&jobPost)

	// validasi
	if jobPost.Title == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Title masih kosong",
			Data:    nil,
		})
	}

	// etc

	var jobDB jobs.Job
	jobDB.Title = jobPost.Title
	jobDB.Category = jobPost.Category
	jobDB.JobDesc = jobPost.JobDesc

	result := configs.DB.Create(&jobDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data Job ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil Post Pekerjaan!",
		Data:    jobDB,
	})
}

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

func GetJobDescController(c echo.Context) error {

	jobdescs := []jobs.JobDesc{}

	result := configs.DB.Find(&jobdescs)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data jobdescs dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data jobdescs",
		Data:    jobdescs,
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

	job := jobs.Job{}

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

func DeleteAllJobsController(c echo.Context) error {

	jobs := []jobs.Job{}
	result := configs.DB.Find(&jobs)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data di DB",
			Data:    nil,
		})
	}

	result = configs.DB.Unscoped().Delete(&jobs)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data Job dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil menghapus data Job berikut",
		Data:    jobs,
	})
}
