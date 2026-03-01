package types

import (
	"api-karang-waru/internal/domain"
	"api-karang-waru/pkg/utils"
)

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func UserResponseFromModel(user *domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Password:  user.Password,
		CreatedAt: utils.FormatTimeHuman(user.CreatedAt),
		UpdatedAt: utils.FormatTimeHuman(user.UpdatedAt),
	}
}
