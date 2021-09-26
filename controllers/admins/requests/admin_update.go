package requests

import "main.go/business/admins"

type AdminUpdate struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Password    string `json:"password"`
	CompanyName string `json:"companyName"`
}

func (admin *AdminUpdate) ToDomain() admins.Domain {
	return admins.Domain{
		Name:        admin.Name,
		Email:       admin.Email,
		Address:     admin.Address,
		Password:    admin.Password,
		CompanyName: admin.CompanyName,
	}
}
