package requests

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
