package responses

import (
	"time"

	"main.go/business/admins"
)

type AdminResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	CompanyName string    `json:"companyName"`
	Token       string    `json:"token"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromDomain(domain admins.Domain) AdminResponse {
	return AdminResponse{
		Id:          domain.Id,
		Name:        domain.Name,
		Email:       domain.Email,
		Address:     domain.Address,
		CompanyName: domain.CompanyName,
		Token:       domain.Token,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func ListFromDomain(domain []admins.Domain) (response []AdminResponse) {
	for _, admin := range domain {
		response = append(response, FromDomain(admin))
	}
	return
}
