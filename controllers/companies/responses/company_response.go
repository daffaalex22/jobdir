package responses

import (
	"time"

	"main.go/business/companies"
	"main.go/controllers/admins/responses"
)

type CompanyResponse struct {
	Id           int                       `json:"id"`
	Name         string                    `json:"name"`
	Address      string                    `json:"address"`
	Description  string                    `json:"description"`
	IsTopCompany bool                      `json:"isTopCompany"`
	Admins       []responses.AdminResponse `json:"admins"`
	CreatedAt    time.Time                 `json:"createdAt"`
	UpdatedAt    time.Time                 `json:"updatedAt"`
}

func FromDomain(company companies.Domain) CompanyResponse {
	return CompanyResponse{
		Id:           company.Id,
		Name:         company.Name,
		Address:      company.Address,
		Description:  company.Description,
		IsTopCompany: company.IsTopCompany,
		Admins:       responses.ListFromDomain(company.Admins),
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
