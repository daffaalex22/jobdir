package companies

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/companies"
	"main.go/drivers/databases/admins"
	"main.go/drivers/databases/jobs"
)

type Companies struct {
	Id           int    `gorm:"primaryKey"`
	Name         string `gorm:"unique"`
	Address      string
	Description  string
	Admins       []admins.Admins `gorm:"foreignKey:CompanyId"`
	IsTopCompany bool
	Jobs         []jobs.Jobs `gorm:"foreignKey:CompanyId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (company *Companies) ToDomain() companies.Domain {
	return companies.Domain{
		Id:           company.Id,
		Name:         company.Name,
		Address:      company.Address,
		Description:  company.Description,
		IsTopCompany: company.IsTopCompany,
		Jobs:         jobs.ListToDomain(company.Jobs),
		Admins:       admins.ListToDomain(company.Admins),
		CreatedAt:    company.CreatedAt,
		UpdatedAt:    company.UpdatedAt,
	}
}

func ListToDomain(Companies []Companies) (result []companies.Domain) {
	for _, company := range Companies {
		result = append(result, company.ToDomain())
	}
	return
}

func FromDomain(domain companies.Domain) Companies {
	return Companies{
		Id:           domain.Id,
		Name:         domain.Name,
		Address:      domain.Address,
		Admins:       admins.ListFromDomain(domain.Admins),
		Jobs:         jobs.ListFromDomain(domain.Jobs),
		Description:  domain.Description,
		IsTopCompany: domain.IsTopCompany,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
