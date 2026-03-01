package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DemografisHandler struct {
	service usecase.DemografisService
}

func NewDemografisHandler(service usecase.DemografisService) *DemografisHandler {
	return &DemografisHandler{service}
}

func (h *DemografisHandler) CreateDemografis(c *gin.Context) {

	var req types.DemografisRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	demografis, err := h.service.CreateDemografis(&req)

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
		Data:    types.DemografisResponseFromModel(demografis),
	})
}

func (h *DemografisHandler) GetDemografis(c *gin.Context) {
	demografis, err := h.service.GetDemografis()

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
		Data:    types.DemografisResponseFromModel(&demografis),
	})

}

func (h *DemografisHandler) UpdateDemografis(c *gin.Context) {

	var req types.DemografisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	demografis, err := h.service.UpdateDemografis(&req)
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
		Data:    types.DemografisResponseFromModel(demografis),
	})
}

// delete
func (h *DemografisHandler) DeleteDemografis(c *gin.Context) {

	err := h.service.DeleteDemografis()
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
