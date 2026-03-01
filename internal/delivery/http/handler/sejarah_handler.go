package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SejarahHandler struct {
	service usecase.SejarahService
}

func NewSejarahHandler(service usecase.SejarahService) *SejarahHandler {
	return &SejarahHandler{service}
}

func (h *SejarahHandler) CreateSejarah(c *gin.Context) {

	var req types.SejarahRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	sejarah, err := h.service.CreateSejarah(&req)

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
		Data:    types.SejarahResponseFromModel(sejarah),
	})
}

func (h *SejarahHandler) GetSejarah(c *gin.Context) {
	sejarah, err := h.service.GetSejarah()

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
		Data:    types.SejarahResponseFromModel(&sejarah),
	})

}

func (h *SejarahHandler) UpdateSejarah(c *gin.Context) {

	var req types.SejarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	sejarah, err := h.service.UpdateSejarah(&req)
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
		Data:    types.SejarahResponseFromModel(sejarah),
	})
}

// delete
func (h *SejarahHandler) DeleteSejarah(c *gin.Context) {

	err := h.service.DeleteSejarah()
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
		Message: "Sejarah deleted successfully",
		Data:    nil,
	})
}
