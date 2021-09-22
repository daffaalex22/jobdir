package jobs

import (
	"time"

	"gorm.io/gorm"
)

type JobDesc struct {
	ID          int            `gorm:"primaryKey, autoIncrement" json:"id"`
	Description string         `json:"description"`
	JobID       int            `json:"jobId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
