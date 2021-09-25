package categories

import (
	"context"

	"gorm.io/gorm"
	"main.go/business/categories"
)

type MysqlCategoryRepository struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository(conn *gorm.DB) categories.Repository {
	return &MysqlCategoryRepository{
		Conn: conn,
	}
}

func (rep *MysqlCategoryRepository) CreateCategory(ctx context.Context, domain categories.Domain) (categories.Domain, error) {

	category := FromDomain(domain)

	result := rep.Conn.Create(&category)

	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	return category.ToDomain(), nil
}

func (rep *MysqlCategoryRepository) GetCategoryById(ctx context.Context, id int) (categories.Domain, error) {
	var category Categories
	result := rep.Conn.First(&category, "id = ?", id)

	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	return category.ToDomain(), nil
}

func (rep *MysqlCategoryRepository) GetAllCategory(ctx context.Context) ([]categories.Domain, error) {
	var category []Categories
	result := rep.Conn.Find(&category)

	// for _, v := range category {
	// 	v.Jobs = jobs.
	// }

	if result.Error != nil {
		return []categories.Domain{}, result.Error
	}

	return ListToDomain(category), nil
}

// func (rep *MysqlCategoryRepository) UpdateCat(ctx context.Context, domain categories.Domain) (Cats.Domain, error) {
// 	var category Categories
// 	result := rep.Conn.First(&category, "email = ? AND password = ?", domain.Email, domain.Password)

// 	if result.Error != nil {
// 		return Cats.Domain{}, result.Error
// 	}

// 	result = rep.Conn.Model(&Cat).Updates(FromDomain(domain))

// 	if result.Error != nil {
// 		return Cats.Domain{}, result.Error
// 	}

// 	result = rep.Conn.Save(&Cat)

// 	if result.Error != nil {
// 		return Cats.Domain{}, result.Error
// 	}

// 	return Cat.ToDomain(), nil
// }

func (rep *MysqlCategoryRepository) DeleteCategoryById(ctx context.Context, id int) (categories.Domain, error) {
	var category Categories
	result := rep.Conn.Where("id = ?", id).Delete(&category)

	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	return category.ToDomain(), nil
}
