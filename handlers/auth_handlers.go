package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req requests.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
		})
		return
	}

	res, err := h.authService.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.APIResponse{
			Code:    "UNAUTHORIZED",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
		})
		return
	}

	res, err := h.authService.SignUp(req.Email, req.Password, req.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.APIResponse{
			Code:    "UNAUTHORIZED",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
