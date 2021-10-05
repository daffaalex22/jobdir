package requests

import "main.go/business/users"

type UserUpdate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func (user *UserUpdate) ToDomain() users.Domain {
	return users.Domain{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Password: user.Password,
	}
}
