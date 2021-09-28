package requests

import (
	"main.go/business/applications"
)

type ApplicationCreate struct {
	UserId int    `json:"UserId"`
	JobId  int    `json:"jobId"`
	Status string `json:"status"`
}

func (application *ApplicationCreate) ToDomain() applications.Domain {
	return applications.Domain{
		UserId: application.UserId,
		JobId:  application.JobId,
		Status: application.Status,
	}
}

func ListToDomain(applications []ApplicationCreate) (result []applications.Domain) {
	for _, application := range applications {
		result = append(result, application.ToDomain())
	}
	return
}
