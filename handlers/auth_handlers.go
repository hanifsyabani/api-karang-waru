package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"
	"strconv"

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

	c.SetCookie(
		"access_token",     // nama cookie
		res.AccessToken,    // value = token JWT
		res.ExpiresIn*3600, // detik (jam → detik)
		"/",                // path
		"",                 // domain (kosong = current domain)
		true,               // secure (harus HTTPS kalau true)
		true,               // httpOnly (tidak bisa diakses via JS)
	)

	c.SetCookie(
		"user_id",
		strconv.FormatUint(uint64(res.User.ID), 10),
		res.ExpiresIn*3600,
		"/",
		"",
		true,
		true,
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

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "logout success"})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	id, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user_id"})
		return
	}

	user, err := h.service.GetUserByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.APIResponse{
			Code:    "NOT_FOUND",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "OK",
		Message: "User retrieved successfully",
		Data:    responses.UserResponseFromModel(user),
	})
}
