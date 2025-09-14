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
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


func UserResponseFromModel(contact *models.User) UserResponse {
	return UserResponse{
		ID:        contact.ID,
		Name:      contact.Name,
		Email:     contact.Email,
		CreatedAt: helpers.FormatTimeHuman(contact.CreatedAt),
		UpdatedAt: helpers.FormatTimeHuman(contact.UpdatedAt),
	}
}


