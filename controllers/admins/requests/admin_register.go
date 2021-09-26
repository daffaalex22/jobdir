package requests

import "main.go/business/admins"

type AdminRegister struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Password    string `json:"password"`
	CompanyName string `json:"companyName"`
}

func (admin *AdminRegister) ToDomain() admins.Domain {
	return admins.Domain{
		Name:        admin.Name,
		Email:       admin.Email,
		Address:     admin.Address,
		Password:    admin.Password,
		CompanyName: admin.CompanyName,
	}
}
