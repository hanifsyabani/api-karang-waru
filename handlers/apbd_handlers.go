package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApbdHandler struct {
	service services.ApbdService
}

func NewApbdHandler(service services.ApbdService) *ApbdHandler {
	return &ApbdHandler{service}
}

func (h *ApbdHandler) CreateApbd(c *gin.Context) {
	var req requests.ApbdRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	apbd, err := h.service.CreateApbd(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// ketika succes
	c.JSON(http.StatusCreated, responses.APIResponse{
		Code:    "CREATED",
		Message: "APbd created successfully",
		Data:    responses.ApbdResponseFromModel(apbd),
	})
}

// get all
func (h *ApbdHandler) GetApbd(c *gin.Context) {
	apbd, err := h.service.GetAllApbd()

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var beritaResponses []responses.ApbdResponse
	for _, apbd := range apbd {
		beritaResponses = append(beritaResponses, responses.ApbdResponseFromModel(&apbd))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
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
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid apbd ID",
			Data:    nil,
		})
		return
	}

	apbd, err := h.service.GetApbdByID(uint(id))
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
		Message: "apbd retrieved successfully",
		Data:    responses.ApbdResponseFromModel(apbd),
	})
}

func (h *ApbdHandler) UpdateApbd(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid apbd ID",
			Data:    nil,
		})
		return
	}

	var req requests.ApbdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	apbd, err := h.service.UpdateApbd(uint(id), &req)
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
		Message: "User updated successfully",
		Data:    responses.ApbdResponseFromModel(apbd),
	})
}

// delete
func (h *ApbdHandler) DeleteApbd(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid apbd ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeleteApbd(uint(id))
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
		Message: "APbd deleted successfully",
		Data:    nil,
	})
}
