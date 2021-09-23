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

func (rep *MysqlUserRepository) GetUserById(ctx context.Context, id int) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "id = ?", id)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) GetAllUser(ctx context.Context) ([]users.Domain, error) {
	var user []Users
	result := rep.Conn.Find(&user)

	if result.Error != nil {
		return []users.Domain{}, result.Error
	}

	return ListToDomain(user), nil
}

func (rep *MysqlUserRepository) UpdateUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "email = ? AND password = ?", domain.Email, domain.Password)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	result = rep.Conn.Model(&user).Updates(Users{
		Email:    domain.Email,
		Password: domain.Password,
		Name:     domain.Name,
		Address:  domain.Address,
	})

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	result = rep.Conn.Save(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}
