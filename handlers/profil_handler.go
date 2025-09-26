package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfilDesaHandler struct {
	service services.ProfilService
}

func NewProfilDesaHandler(service services.ProfilService) *ProfilDesaHandler {
	return &ProfilDesaHandler{service}
}

func (h *ProfilDesaHandler) CreateProfil(c *gin.Context) {

	var req requests.ProfilDesaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	profil, err := h.service.CreateProfil(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, responses.APIResponse{
		Code:    "CREATED",
		Message: "Profil created successfully",
		Data:    responses.ProfilResponseFromModel(profil),
	})
}

func (h *ProfilDesaHandler) GetProfil(c *gin.Context) {
	profile, err := h.service.GetProfil()

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,	
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "OK",
		Message: "Profil retrieved successfully",
		Data:    responses.ProfilResponseFromModel(&profile),
	})

}

func (h *ProfilDesaHandler) UpdateProfil(c *gin.Context) {

	var req requests.ProfilDesaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	profil, err := h.service.UpdateProfil(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "OK",
		Message: "Profil updated successfully",
		Data:    responses.ProfilResponseFromModel(profil),
	})
}

// delete
func (h *ProfilDesaHandler) DeleteProfil(c *gin.Context) {

	err := h.service.DeleteProfil()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "OK",
		Message: "Profil deleted successfully",
		Data:    nil,
	})
}







