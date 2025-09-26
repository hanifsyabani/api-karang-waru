package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DemografisHandler struct {
	service services.DemografisService
}

func NewDemografisHandler(service services.DemografisService) *DemografisHandler {
	return &DemografisHandler{service}
}


func (h *DemografisHandler) CreateDemografis(c *gin.Context) {

	var req requests.DemografisRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	demografis, err := h.service.CreateDemografis(&req)

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
		Data:    responses.DemografisResponseFromModel(demografis),
	})
}

func (h *DemografisHandler) GetDemografis(c *gin.Context) {
	demografis, err := h.service.GetDemografis()

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
		Data:    responses.DemografisResponseFromModel(&demografis),
	})

}

func (h *DemografisHandler) UpdateDemografis(c *gin.Context) {

	var req requests.DemografisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	demografis, err := h.service.UpdateDemografis(&req)
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
		Data:    responses.DemografisResponseFromModel(demografis),
	})
}

// delete
func (h *DemografisHandler) DeleteDemografis(c *gin.Context) {

	err := h.service.DeleteDemografis()
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