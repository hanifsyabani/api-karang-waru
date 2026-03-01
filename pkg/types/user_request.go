package types

type UserRequest struct {
	Name     string `json:"name" binding:"required,max=255"`
	Email    string `json:"email" binding:"required,max=255"`
	Role     string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

