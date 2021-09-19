package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Address   string         `json:"address"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
