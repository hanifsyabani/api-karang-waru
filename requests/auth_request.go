package requests


type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

