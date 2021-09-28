package requests

import (
	"main.go/business/companies"
)

type CompanyUpdate struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	IsTopCompany bool   `json:"isTopCompany"`
}

func (company *CompanyUpdate) ToDomain() companies.Domain {
	return companies.Domain{
		Name:         company.Name,
		Address:      company.Address,
		Description:  company.Description,
		IsTopCompany: company.IsTopCompany,
	}
}
