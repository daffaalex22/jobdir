package companies

import (
	"context"

	"gorm.io/gorm"
	"main.go/business/companies"
)

type MysqlCompanyRepository struct {
	Conn *gorm.DB
}

func NewMysqlCompanyRepository(conn *gorm.DB) companies.Repository {
	return &MysqlCompanyRepository{
		Conn: conn,
	}
}

func (rep *MysqlCompanyRepository) GetCompanyById(ctx context.Context, id int) (companies.Domain, error) {
	var company Companies
	result := rep.Conn.Preload("Admins").Preload("Jobs.Category").Preload("Jobs.Applications").First(&company, "id = ?", id)

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	return company.ToDomain(), nil
}

func (rep *MysqlCompanyRepository) GetAllCompany(ctx context.Context) ([]companies.Domain, error) {
	var company []Companies

	result := rep.Conn.Find(&company)
	if result.Error != nil {
		return []companies.Domain{}, result.Error
	}

	return ListToDomain(company), nil
}

func (rep *MysqlCompanyRepository) UpdateCompany(ctx context.Context, domain companies.Domain) (companies.Domain, error) {
	var company Companies

	result := rep.Conn.Where(FromDomain(domain)).First(&company)

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	result = rep.Conn.Model(&company).Updates(FromDomain(domain))

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	result = rep.Conn.Save(&company)

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	return company.ToDomain(), nil
}

func (rep *MysqlCompanyRepository) DeleteCompany(ctx context.Context, id int) (companies.Domain, error) {
	var company Companies
	result := rep.Conn.Where("id = ?", id).Delete(&company)

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	return company.ToDomain(), nil
}

func (rep *MysqlCompanyRepository) RegisterCompany(ctx context.Context, domain companies.Domain) (companies.Domain, error) {
	company := FromDomain(domain)

	result := rep.Conn.Create(&company)

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	return company.ToDomain(), nil
}
