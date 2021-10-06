package responses

import (
	"time"

	"main.go/business/admins"
	"main.go/controllers/jobs/responses"
)

type AdminResponse struct {
	Id          int                     `json:"id"`
	Name        string                  `json:"name"`
	Email       string                  `json:"email"`
	Address     string                  `json:"address"`
	CompanyId   int                     `json:"companyId"`
	JobsCreated []responses.JobResponse `json:"jobsCreated"`
	// Token       string                  `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoginAdminResponse struct {
	Id          int                     `json:"id"`
	Name        string                  `json:"name"`
	Email       string                  `json:"email"`
	Address     string                  `json:"address"`
	CompanyId   int                     `json:"companyId"`
	JobsCreated []responses.JobResponse `json:"jobsCreated"`
	Token       string                  `json:"token"`
	CreatedAt   time.Time               `json:"createdAt"`
	UpdatedAt   time.Time               `json:"updatedAt"`
}

func FromDomain(domain admins.Domain) AdminResponse {
	return AdminResponse{
		Id:          domain.Id,
		Name:        domain.Name,
		Email:       domain.Email,
		Address:     domain.Address,
		CompanyId:   domain.CompanyId,
		JobsCreated: responses.ListFromDomain(domain.JobsCreated),
		// Token:       domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ListFromDomain(domain []admins.Domain) (response []AdminResponse) {
	for _, admin := range domain {
		response = append(response, FromDomain(admin))
	}
	return
}

func FromDomainLogin(domain admins.Domain) LoginAdminResponse {
	return LoginAdminResponse{
		Id:          domain.Id,
		Name:        domain.Name,
		Email:       domain.Email,
		Address:     domain.Address,
		CompanyId:   domain.CompanyId,
		JobsCreated: responses.ListFromDomain(domain.JobsCreated),
		Token:       domain.Token,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func ListFromDomainLogin(domain []admins.Domain) (response []LoginAdminResponse) {
	for _, admin := range domain {
		response = append(response, FromDomainLogin(admin))
	}
	return
}
