package responses

type SignInResponse struct {
	AccessToken string       `json:"access_token"`
	TokenType   string       `json:"token_type"`
	ExpiresIn   int          `json:"expires_in_h"`
	User        UserResponse `json:"user"`
}

type SignUpResponse struct {
	User UserResponse `json:"user"`
}
 