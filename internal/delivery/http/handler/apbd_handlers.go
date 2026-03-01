package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApbdHandler struct {
	service usecase.ApbdService
}

func NewApbdHandler(service usecase.ApbdService) *ApbdHandler {
	return &ApbdHandler{service}
}

func (h *ApbdHandler) CreateApbd(c *gin.Context) {
	var req types.ApbdRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	apbd, err := h.service.CreateApbd(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// ketika succes
	c.JSON(http.StatusCreated, types.APIResponse{
		Code:    "CREATED",
		Message: "APbd created successfully",
		Data:    types.ApbdResponseFromModel(apbd),
	})
}

// get all
func (h *ApbdHandler) GetApbd(c *gin.Context) {
	apbd, err := h.service.GetAllApbd()

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var beritaResponses []types.ApbdResponse
	for _, apbd := range apbd {
		beritaResponses = append(beritaResponses, types.ApbdResponseFromModel(&apbd))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "APbd retrieved successfully",
		Data:    beritaResponses,
	})

}

// get by id
func (h *ApbdHandler) GetApbdByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid apbd ID",
			Data:    nil,
		})
		return
	}

	apbd, err := h.service.GetApbdByID(uint(id))
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
		Message: "apbd retrieved successfully",
		Data:    types.ApbdResponseFromModel(apbd),
	})
}

func (h *ApbdHandler) UpdateApbd(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid apbd ID",
			Data:    nil,
		})
		return
	}

	var req types.ApbdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	apbd, err := h.service.UpdateApbd(uint(id), &req)
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
		Message: "User updated successfully",
		Data:    types.ApbdResponseFromModel(apbd),
	})
}

// delete
func (h *ApbdHandler) DeleteApbd(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid apbd ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeleteApbd(uint(id))
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
		Message: "APbd deleted successfully",
		Data:    nil,
	})
}

