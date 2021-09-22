package jobs

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int            `gorm:"primaryKey, autoIncrement" json:"id"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
