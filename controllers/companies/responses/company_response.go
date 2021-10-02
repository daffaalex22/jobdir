package responses

import (
	"time"

	"main.go/business/companies"
	_adminResponse "main.go/controllers/admins/responses"
	_jobResponse "main.go/controllers/jobs/responses"
)

type CompanyResponse struct {
	Id           int                            `json:"id"`
	Name         string                         `json:"name"`
	Address      string                         `json:"address"`
	Description  string                         `json:"description"`
	IsTopCompany bool                           `json:"isTopCompany"`
	Admins       []_adminResponse.AdminResponse `json:"admins"`
	Jobs         []_jobResponse.JobResponse     `json:"jobs"`
	CreatedAt    time.Time                      `json:"createdAt"`
	UpdatedAt    time.Time                      `json:"updatedAt"`
}

type CompanyAllResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	IsTopCompany bool   `json:"isTopCompany"`
	// Admins       []_adminResponse.AdminResponse `json:"admins"`
	// Jobs         []_jobResponse.JobResponse     `json:"jobs"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(company companies.Domain) CompanyResponse {
	return CompanyResponse{
		Id:           company.Id,
		Name:         company.Name,
		Address:      company.Address,
		Description:  company.Description,
		IsTopCompany: company.IsTopCompany,
		Admins:       _adminResponse.ListFromDomain(company.Admins),
		Jobs:         _jobResponse.ListFromDomain(company.Jobs),
		CreatedAt:    company.CreatedAt,
		UpdatedAt:    company.UpdatedAt,
	}
}

func ListFromDomain(domain []companies.Domain) (response []CompanyResponse) {
	for _, admin := range domain {
		response = append(response, FromDomain(admin))
	}
	return
}

func FromDomainAll(company companies.Domain) CompanyAllResponse {
	return CompanyAllResponse{
		Id:           company.Id,
		Name:         company.Name,
		Address:      company.Address,
		Description:  company.Description,
		IsTopCompany: company.IsTopCompany,
		// Admins:       _adminResponse.ListFromDomain(company.Admins),
		// Jobs:         _jobResponse.ListFromDomain(company.Jobs),
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
	}
}

func ListFromDomainAll(domain []companies.Domain) (response []CompanyAllResponse) {
	for _, admin := range domain {
		response = append(response, FromDomainAll(admin))
	}
	return
}
