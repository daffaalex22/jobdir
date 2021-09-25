package categories

import (
	"time"

	"gorm.io/gorm"
	"main.go/business/categories"
	"main.go/drivers/databases/jobs"
)

type Categories struct {
	Id       int    `gorm:"primaryKey"`
	Category string `gorm:"unique"`
	// CategoryDescription string
	Jobs      []jobs.Jobs `gorm:"foreignKey:CategoryId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Categories *Categories) ToDomain() categories.Domain {
	return categories.Domain{
		Id:       Categories.Id,
		Category: Categories.Category,
		Jobs:     jobs.ListToDomain(Categories.Jobs),
		// CategoryDescription: Categories.CategoryDescription,
		CreatedAt: Categories.CreatedAt,
		UpdatedAt: Categories.UpdatedAt,
	}
}

func ListToDomain(Categories []Categories) (result []categories.Domain) {
	for _, category := range Categories {
		result = append(result, category.ToDomain())
	}
	return
}

func FromDomain(domain categories.Domain) Categories {
	return Categories{
		Id:       domain.Id,
		Category: domain.Category,
		Jobs:     jobs.ListFromDomain(domain.Jobs),
		// CategoryDescription: domain.CategoryDescription,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
