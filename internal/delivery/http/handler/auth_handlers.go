package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"api-karang-waru/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *usecase.AuthService
}

func NewAuthHandler(authService *usecase.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req types.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
		})
		return
	}

	res, err := h.authService.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, types.APIResponse{
			Code:    "UNAUTHORIZED",
			Message: err.Error(),
		})
		return
	}

	utils.SetAuthCookies(
		c,
		res.AccessToken,
		uint(res.ExpiresIn),
	)

	c.JSON(http.StatusOK, res)
}
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
		})
		return
	}

	res, err := h.authService.SignUp(req.Email, req.Password, req.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, types.APIResponse{
			Code:    "UNAUTHORIZED",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	utils.ClearAuthCookies(c)
	c.JSON(200, gin.H{"message": "logout success"})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists  {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "unauthorized",
			Data:    nil,
		})
		return
	}

	user, err := h.service.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, types.APIResponse{
			Code:    "NOT_FOUND",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "User retrieved successfully",
		Data:    types.UserResponseFromModel(user),
	})
}

