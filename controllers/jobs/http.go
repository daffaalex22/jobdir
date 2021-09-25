package jobs

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"main.go/business/jobs"
	"main.go/controllers"
	"main.go/controllers/jobs/requests"
	"main.go/controllers/jobs/responses"
)

type JobController struct {
	JobUseCase jobs.Usecase
}

func NewJobController(JobUseCase jobs.Usecase) *JobController {
	return &JobController{
		JobUseCase: JobUseCase,
	}
}

// func (JobController JobController) FillJobs(c echo.Context, category []jobs.CategoryDomain) error {

// 	ctx := c.Request().Context()
// 	res, err := JobController.JobUseCase.FillJobs(ctx, category)

// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.ToListCategoryDomain(res))
// }

func (JobController JobController) CreateJob(c echo.Context) error {
	fmt.Println("CreateJob")
	jobCreate := requests.JobCreate{}
	c.Bind(&jobCreate)

	ctx := c.Request().Context()
	job, err := JobController.JobUseCase.CreateJob(ctx, jobCreate.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(job))
}

func (JobController JobController) GetAllJobs(c echo.Context) error {
	fmt.Println("GetAllJobs")

	ctx := c.Request().Context()
	job, err := JobController.JobUseCase.GetAllJobs(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ListFromDomain(job))
}

func (JobController JobController) GetJobById(c echo.Context) error {
	fmt.Println("GetJobById")

	jobId, err := strconv.Atoi(c.Param("jobId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	job, err := JobController.JobUseCase.GetJobById(ctx, jobId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(job))
}

func (JobController JobController) DeleteAllJobs(c echo.Context) error {
	fmt.Println("HardDeleteAllJobRecords")

	ctx := c.Request().Context()
	error := JobController.JobUseCase.DeleteAllJobs(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(jobs.Domain{}))
}

func (JobController JobController) DeleteJobById(c echo.Context) error {
	fmt.Println("DeleteJobById")

	jobId, err := strconv.Atoi(c.Param("jobId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	job, err := JobController.JobUseCase.DeleteJobById(ctx, jobId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(job))
}

func (JobController JobController) SearchJobs(c echo.Context) error {
	fmt.Println("GetAllJobs")

	title := c.QueryParam("title")

	ctx := c.Request().Context()
	job, err := JobController.JobUseCase.SearchJobs(ctx, title)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ListFromDomain(job))
}
