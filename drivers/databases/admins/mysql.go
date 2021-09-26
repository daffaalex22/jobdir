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

func (rep *MysqlAdminRepository) Login(ctx context.Context, email string, password string) (admins.Domain, error) {
	var admin Admins

	result := rep.Conn.First(&admin, "email = ?", email)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	err := encrypt.CheckPassword(password, admin.Password)

	if err != nil {
		return admins.Domain{}, result.Error
	}

	return admin.ToDomain(), nil
}

func (rep *MysqlAdminRepository) GetAdminById(ctx context.Context, id int) (admins.Domain, error) {
	var admin Admins
	result := rep.Conn.First(&admin, "id = ?", id)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return admin.ToDomain(), nil
}

func (rep *MysqlAdminRepository) GetAllAdmin(ctx context.Context) ([]admins.Domain, error) {
	var Admin []Admins
	result := rep.Conn.Find(&Admin)

	if result.Error != nil {
		return []admins.Domain{}, result.Error
	}

	return ListToDomain(Admin), nil
}

func (rep *MysqlAdminRepository) UpdateAdmin(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	var Admin Admins

	result := rep.Conn.First(&Admin, "email = ? AND password = ?", domain.Email, domain.Password)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	result = rep.Conn.Model(&Admin).Updates(FromDomain(domain))

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	result = rep.Conn.Save(&Admin)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return Admin.ToDomain(), nil
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

func (rep *MysqlAdminRepository) HardDeleteAllAdmins(ctx context.Context) error {
	var Admin []Admins
	result := rep.Conn.Find(&Admin)

	if result.Error != nil {
		return result.Error
	}

	result = rep.Conn.Unscoped().Delete(&Admin)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
