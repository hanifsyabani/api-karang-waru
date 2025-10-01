package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SejarahHandler struct {
	service services.SejarahService
}

func NewSejarahHandler(service services.SejarahService) *SejarahHandler {
	return &SejarahHandler{service}
}


func (h *SejarahHandler) CreateSejarah(c *gin.Context) {

	var req requests.SejarahRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	sejarah, err := h.service.CreateSejarah(&req)

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
		Data:    responses.SejarahResponseFromModel(sejarah),
	})
}

func (h *SejarahHandler) GetSejarah(c *gin.Context) {
	sejarah, err := h.service.GetSejarah()

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
		Data:    responses.SejarahResponseFromModel(&sejarah),
	})

}

func (h *SejarahHandler) UpdateSejarah(c *gin.Context) {

	var req requests.SejarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	sejarah, err := h.service.UpdateSejarah(&req)
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
		Data:    responses.SejarahResponseFromModel(sejarah),
	})
}

// delete
func (h *SejarahHandler) DeleteSejarah(c *gin.Context) {

	err := h.service.DeleteSejarah()
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
		Message: "Sejarah deleted successfully",
		Data:    nil,
	})
}