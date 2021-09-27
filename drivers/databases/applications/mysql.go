package applications

import (
	"context"

	"gorm.io/gorm"
	"main.go/business/applications"
)

type MysqlApplicationRepository struct {
	Conn *gorm.DB
}

func NewMysqlApplicationRepository(conn *gorm.DB) applications.Repository {
	return &MysqlApplicationRepository{
		Conn: conn,
	}
}

// func (rep *MysqlApplicationRepository) FillApplications(ctx context.Context, categories []applications.CategoryDomain) ([]Applications.CategoryDomain, error) {

// 	for _, category := range categories {
// 		var Application []Applications
// 		result := rep.Conn.Where("categoryId = ?", category.Id).Find(&Application)
// 		if result.Error != nil {
// 			return categories, result.Error
// 		}
// 		category.Applications = append(category.Applications, ListToDomain(Application)...)
// 	}
// 	return categories, nil
// }

func (rep *MysqlApplicationRepository) CreateApplication(ctx context.Context, domain applications.Domain) (applications.Domain, error) {
	application := FromDomain(domain)

	result := rep.Conn.Create(&application)

	if result.Error != nil {
		return applications.Domain{}, result.Error
	}

	return application.ToDomain(), nil
}

func (rep *MysqlApplicationRepository) GetAllApplications(ctx context.Context) ([]applications.Domain, error) {
	var Application []Applications
	result := rep.Conn.Find(&Application)

	if result.Error != nil {
		return []applications.Domain{}, result.Error
	}

	return ListToDomain(Application), nil
}

// func (rep *MysqlApplicationRepository) GetApplicationById(ctx context.Context, id int) (applications.Domain, error) {
// 	var Application Applications
// 	result := rep.Conn.First(&Application, "id = ?", id)

// 	if result.Error != nil {
// 		return applications.Domain{}, result.Error
// 	}

// 	return Application.ToDomain(), nil
// }

func (rep *MysqlApplicationRepository) DeleteAllApplications(ctx context.Context) error {
	var Applications []Applications

	result := rep.Conn.Find(&Applications)
	if result.Error != nil {
		return result.Error
	}

	result = rep.Conn.Unscoped().Delete(&Applications)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func (rep *MysqlApplicationRepository) DeleteApplicationById(ctx context.Context, id int) (applications.Domain, error) {
// 	var Application Applications
// 	result := rep.Conn.Where("id = ?", id).Delete(&Application)

// 	if result.Error != nil {
// 		return applications.Domain{}, result.Error
// 	}

// 	return Application.ToDomain(), nil
// }

// func (rep *MysqlApplicationRepository) SearchApplications(ctx context.Context, title string) ([]applications.Domain, error) {
// 	var Application []Applications
// 	result := rep.Conn.Where("title LIKE ?", title+"%").Find(&Application)
// 	if result.Error != nil {
// 		return []applications.Domain{}, result.Error
// 	}

// 	result = rep.Conn.Where("title LIKE ?", title+"%").Find(&Application)
// 	if result.Error != nil {
// 		return []applications.Domain{}, result.Error
// 	}

// 	result = rep.Conn.Where("title LIKE ?", "%"+title+"%").Find(&Application)
// 	if result.Error != nil {
// 		return []applications.Domain{}, result.Error
// 	}

// 	return ListToDomain(Application), nil
// }

// func (rep *MysqlApplicationRepository) FilterApplicationByCategory(ctx context.Context, categoryId int) ([]applications.Domain, error) {
// 	var Application []Applications
// 	result := rep.Conn.Where("category_id = ?", categoryId).Find(&Application)
// 	if result.Error != nil {
// 		return []applications.Domain{}, result.Error
// 	}

// 	return ListToDomain(Application), nil
// }
