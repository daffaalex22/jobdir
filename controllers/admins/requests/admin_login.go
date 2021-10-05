package requests

import "main.go/business/admins"

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *AdminLogin) ToDomain() admins.Domain {
	return admins.Domain{
		Email:    user.Email,
		Password: user.Password,
	}
}
