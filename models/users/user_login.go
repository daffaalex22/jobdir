package users

type UserLogin struct {
	Email    string `json:"email"`
	Password int    `json:"password"`
}
