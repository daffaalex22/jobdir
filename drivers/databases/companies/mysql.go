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
	var Company Companies
	result := rep.Conn.Preload("Admins.JobsCreated.Category").Preload("Jobs.Category").Preload("Jobs.Applications").First(&Company, "id = ?", id)

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	return Company.ToDomain(), nil
}

func (rep *MysqlCompanyRepository) GetAllCompany(ctx context.Context) ([]companies.Domain, error) {
	var Company []Companies

	result := rep.Conn.Find(&Company)
	if result.Error != nil {
		return []companies.Domain{}, result.Error
	}

	return ListToDomain(Company), nil
}

func (rep *MysqlCompanyRepository) UpdateCompany(ctx context.Context, domain companies.Domain) (companies.Domain, error) {
	var Company Companies

	// result := rep.Conn.First(&Company, "email = ? AND password = ?", domain.Email, domain.Password)

	// if result.Error != nil {
	// 	return companies.Domain{}, result.Error
	// }

	result := rep.Conn.Model(&Company).Updates(FromDomain(domain))

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	result = rep.Conn.Save(&Company)

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	return Company.ToDomain(), nil
}

func (rep *MysqlCompanyRepository) DeleteCompany(ctx context.Context, id int) (companies.Domain, error) {
	var Company Companies
	result := rep.Conn.Where("id = ?", id).Delete(&Company)

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	return Company.ToDomain(), nil
}

func (rep *MysqlCompanyRepository) RegisterCompany(ctx context.Context, domain companies.Domain) (companies.Domain, error) {
	Company := FromDomain(domain)

	result := rep.Conn.Create(&Company)

	if result.Error != nil {
		return companies.Domain{}, result.Error
	}

	return Company.ToDomain(), nil
}

func (rep *MysqlCompanyRepository) HardDeleteAllCompanies(ctx context.Context) error {
	var Company []Companies
	result := rep.Conn.Find(&Company)

	if result.Error != nil {
		return result.Error
	}

	result = rep.Conn.Unscoped().Delete(&Company)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
