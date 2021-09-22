package jobs

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID        int            `gorm:"primaryKey, autoIncrement" json:"id"`
	Title     string         `json:"title"`
	Category  string         `json:"category"`
	JobDesc   JobDesc        `json:"jobDesc"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
