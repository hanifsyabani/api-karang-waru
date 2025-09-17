package services

import (
	"api-karang-waru/config"
	"api-karang-waru/helpers"
	"api-karang-waru/models"
	"api-karang-waru/responses"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) SignIn(email, password string) (*responses.SignInResponse, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	hashed := helpers.CheckPasswordHash(password, user.Password)
	if !hashed {
		return nil, fmt.Errorf("invalid email or password")
	}

	secret := config.GetEnv("JWT_SECRET", "")
	if secret == "" {
		return nil, fmt.Errorf("jwt secret not set")
	}

	expHours := 72
	if v := config.GetEnv("JWT_EXPIRE_HOURS", ""); v != "" {
		if p, err := strconv.Atoi(v); err == nil {
			expHours = p
		}
	}

	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(expHours)).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}

	return &responses.SignInResponse{
		AccessToken: ss,
		TokenType:   "Bearer",
		ExpiresIn:   expHours,
		User : responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
			CreatedAt: helpers.FormatTimeHuman(user.CreatedAt),
			UpdatedAt: helpers.FormatTimeHuman(user.UpdatedAt),
		},
	}, nil

}

func (s *AuthService) SignUp(email, password, name string) (*responses.SignUpResponse, error) {
	var existingUser models.User

	if err := config.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("user already exists")
	}

	hashed, err := helpers.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{Name: name, Email: email, Password: hashed}
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &responses.SignUpResponse{
		User: responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			CreatedAt: helpers.FormatTimeHuman(user.CreatedAt),
			UpdatedAt: helpers.FormatTimeHuman(user.UpdatedAt),
		},
	}, nil
}
