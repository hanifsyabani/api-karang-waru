package usecase

import (
	"api-karang-waru/config"
	"api-karang-waru/pkg/utils"
	"api-karang-waru/internal/domain"
	"api-karang-waru/pkg/types"
	"fmt"
	"strconv"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	repository UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) SignIn(email, password string) (*types.SignInResponse, error) {
	var user domain.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	hashed := utils.CheckPasswordHash(password, user.Password)
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

	return &types.SignInResponse{
		AccessToken: ss,
		TokenType:   "Bearer",
		ExpiresIn:   expHours,
		User : types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
			CreatedAt: utils.FormatTimeHuman(user.CreatedAt),
			UpdatedAt: utils.FormatTimeHuman(user.UpdatedAt),
		},
	}, nil

}

func (s *AuthService) SignUp(email, password, name string) (*types.SignUpResponse, error) {
	var existingUser domain.User

	if err := config.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("user already exists")
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := domain.User{Name: name, Email: email, Password: hashed}
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &types.SignUpResponse{
		User: types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			CreatedAt: utils.FormatTimeHuman(user.CreatedAt),
			UpdatedAt: utils.FormatTimeHuman(user.UpdatedAt),
		},
	}, nil
}


func (s *AuthService) GetProfile(id uint) (*types.UserResponse, error) {

	user,err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}
	return &types.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		CreatedAt: utils.FormatTimeHuman(user.CreatedAt),
		UpdatedAt: utils.FormatTimeHuman(user.UpdatedAt),
	}, nil
}

