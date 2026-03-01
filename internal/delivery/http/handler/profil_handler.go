package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfilDesaHandler struct {
	service usecase.ProfilService
}

func NewProfilDesaHandler(service usecase.ProfilService) *ProfilDesaHandler {
	return &ProfilDesaHandler{service}
}

func (h *ProfilDesaHandler) CreateProfil(c *gin.Context) {

	var req types.ProfilDesaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	profil, err := h.service.CreateProfil(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, types.APIResponse{
		Code:    "CREATED",
		Message: "Profil created successfully",
		Data:    types.ProfilResponseFromModel(profil),
	})
}

func (h *ProfilDesaHandler) GetProfil(c *gin.Context) {
	profile, err := h.service.GetProfil()

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,	
		})
		return
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Profil retrieved successfully",
		Data:    types.ProfilResponseFromModel(&profile),
	})

}

func (h *ProfilDesaHandler) UpdateProfil(c *gin.Context) {

	var req types.ProfilDesaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	profil, err := h.service.UpdateProfil(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Profil updated successfully",
		Data:    types.ProfilResponseFromModel(profil),
	})
}

// delete
func (h *ProfilDesaHandler) DeleteProfil(c *gin.Context) {

	err := h.service.DeleteProfil()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Profil deleted successfully",
		Data:    nil,
	})
}

