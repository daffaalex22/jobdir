package requests

import (
	"main.go/business/companies"
)

type CompanyCreate struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	IsTopCompany bool   `json:"isTopCompany"`
}

func (admin *CompanyCreate) ToDomain() companies.Domain {
	return companies.Domain{
		Name:         admin.Name,
		Address:      admin.Address,
		Description:  admin.Description,
		IsTopCompany: admin.IsTopCompany,
	}
}
