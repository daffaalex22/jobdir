package responses

import (
	"time"

	"main.go/business/users"
	"main.go/controllers/applications/responses"
)

type UserResponse struct {
	Id           int                             `json:"id"`
	Name         string                          `json:"name"`
	Email        string                          `json:"email"`
	Address      string                          `json:"address"`
	Token        string                          `json:"token"`
	Applications []responses.ApplicationResponse `json:"applications"`
	CreatedAt    time.Time                       `json:"createdAt"`
	UpdatedAt    time.Time                       `json:"updatedAt"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:           domain.Id,
		Name:         domain.Name,
		Email:        domain.Email,
		Address:      domain.Address,
		Applications: responses.ListFromDomain(domain.Applications),
		Token:        domain.Token,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func ListFromDomain(domain []users.Domain) (response []UserResponse) {
	for _, user := range domain {
		response = append(response, FromDomain(user))
	}
	return
}
