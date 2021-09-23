package requests

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
