package services

import (
	"api-karang-waru/config"
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthService struct {
	BaseURL string
	APIKey  string
}

func NewAuthService() *AuthService {
	return &AuthService{
		BaseURL: config.GetEnv("SUPABASE_URL", ""),
		APIKey:  config.GetEnv("SUPABASE_API_KEY", ""),
	}
}

func (s *AuthService) SignIn(email, password string) (*responses.SignInResponse, error) {
	url := fmt.Sprintf("%s/auth/v1/token?grant_type=password", s.BaseURL)

	body, _ := json.Marshal(requests.SignInRequest{
		Email:    email,
		Password: password,
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("apikey", s.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result responses.SignInResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.AccessToken == "" {
		return nil, fmt.Errorf("login gagal: %v", resp.Status)
	}

	return &result, nil
}

func (s *AuthService) SignUp(email, password string) (*responses.SignUpResponse, error) {
	url := fmt.Sprintf("%s/auth/v1/signup", s.BaseURL)

	body, _ := json.Marshal(map[string] interface{}{
		"email":    email,
		"password": password,
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("apikey", s.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result responses.SignUpResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
