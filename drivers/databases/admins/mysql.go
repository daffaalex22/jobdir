package admins

import (
	"context"

	"gorm.io/gorm"
	"main.go/business/admins"
	"main.go/helpers/encrypt"
)

type MysqlAdminRepository struct {
	Conn *gorm.DB
}

func NewMysqlAdminRepository(conn *gorm.DB) admins.Repository {
	return &MysqlAdminRepository{
		Conn: conn,
	}
}

func (rep *MysqlAdminRepository) Login(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	var admin Admins

	result := rep.Conn.Preload("JobsCreated.Category").Preload("JobsCreated.Applications").First(&admin, "email = ?", domain.Email)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	err := encrypt.CheckPassword(domain.Password, admin.Password)

	if err != nil {
		return admins.Domain{}, result.Error
	}

	return admin.ToDomain(), nil
}

func (rep *MysqlAdminRepository) GetAdminById(ctx context.Context, id int) (admins.Domain, error) {
	var admin Admins
	result := rep.Conn.Preload("JobsCreated.Category").Preload("JobsCreated.Applications").First(&admin, "id = ?", id)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return admin.ToDomain(), nil
}

func (rep *MysqlAdminRepository) GetAllAdmin(ctx context.Context) ([]admins.Domain, error) {
	var Admin []Admins
	result := rep.Conn.Preload("JobsCreated.Category").Preload("JobsCreated.Applications").Find(&Admin)

	if result.Error != nil {
		return []admins.Domain{}, result.Error
	}

	return ListToDomain(Admin), nil
}

func (rep *MysqlAdminRepository) UpdateAdmin(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	var admin Admins

	result := rep.Conn.Preload("JobsCreated.Category").Preload("JobsCreated.Applications").First(&admin, "email = ?", domain.Email)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	err := encrypt.CheckPassword(domain.Password, admin.Password)

	if err != nil {
		return admins.Domain{}, result.Error
	}

	hashedPassword, err := encrypt.Hash(domain.Password)
	if err != nil {
		return admins.Domain{}, err
	}

	domain.Password = hashedPassword
	result = rep.Conn.Model(&admin).Updates(FromDomain(domain))

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	result = rep.Conn.Save(&admin)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return admin.ToDomain(), nil
}

func (rep *MysqlAdminRepository) DeleteAdmin(ctx context.Context, id int) (admins.Domain, error) {
	var Admin Admins
	result := rep.Conn.Where("id = ?", id).Delete(&Admin)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return Admin.ToDomain(), nil
}

func (rep *MysqlAdminRepository) RegisterAdmin(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	admin := FromDomain(domain)

	hashedPassword, err := encrypt.Hash(domain.Password)
	if err != nil {
		return admins.Domain{}, err
	}

	admin.Password = hashedPassword

	result := rep.Conn.Create(&admin)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return admin.ToDomain(), nil
}
