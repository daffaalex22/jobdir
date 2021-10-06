package users

import (
	"context"

	"gorm.io/gorm"
	"main.go/business/users"
	"main.go/helpers/encrypt"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user Users

	result := rep.Conn.Preload("Applications").First(&user, "email = ?", domain.Email)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	err := encrypt.CheckPassword(domain.Password, user.Password)

	if err != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) GetUserById(ctx context.Context, id int) (users.Domain, error) {
	var user Users
	result := rep.Conn.Preload("Applications").First(&user, "id = ?", id)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) GetAllUser(ctx context.Context) ([]users.Domain, error) {
	var user []Users
	result := rep.Conn.Preload("Applications").Find(&user)

	if result.Error != nil {
		return []users.Domain{}, result.Error
	}

	return ListToDomain(user), nil
}

func (rep *MysqlUserRepository) UpdateUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user Users

	result := rep.Conn.First(&user, "email = ?", domain.Email)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	err := encrypt.CheckPassword(domain.Password, user.Password)

	if err != nil {
		return users.Domain{}, result.Error
	}

	result = rep.Conn.Model(&user).Updates(FromDomain(domain))

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	result = rep.Conn.Save(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) DeleteUser(ctx context.Context, id int) (users.Domain, error) {
	var user Users
	result := rep.Conn.Where("id = ?", id).Delete(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) RegisterUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	user := FromDomain(domain)

	hashedPassword, err := encrypt.Hash(domain.Password)
	if err != nil {
		return users.Domain{}, err
	}

	user.Password = hashedPassword

	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}
