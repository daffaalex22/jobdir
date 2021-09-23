package users

import (
	"context"

	"gorm.io/gorm"
	"main.go/business/users"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "email = ? AND password = ?", email, password)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) GetById(ctx context.Context, id int) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "id = ?", id)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}
