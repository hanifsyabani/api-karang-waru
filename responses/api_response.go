package responses

import (
	"api-karang-waru/helpers"
	"api-karang-waru/models"
)

type APIResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func UserResponseFromModel(user *models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: helpers.FormatTimeHuman(user.CreatedAt),
		UpdatedAt: helpers.FormatTimeHuman(user.UpdatedAt),
	}
}
