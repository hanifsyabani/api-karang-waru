package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PendudukHandler struct {
	service services.PendudukService
}

func NewPendudukHandler(service services.PendudukService) *PendudukHandler {
	return &PendudukHandler{service}
}

func (h *PendudukHandler) CreatePenduduk(c *gin.Context) {
	var req requests.PendudukRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	penduduk, err := h.service.CreatePenduduk(&req)

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
		Message: "penduduk created successfully",
		Data:    responses.PendudukResponseFromModel(penduduk),
	})
}

// get all
func (h *PendudukHandler) GetAllPenduduk(c *gin.Context) {
	search := c.Query("query")
	sortBy := c.DefaultQuery("sortBy", "created_at")
	sortOrder := c.DefaultQuery("sortOrder", "desc")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	penduduk, err := h.service.GetAllPenduduk(search, page, limit, sortBy, sortOrder)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var pendudukResponse []responses.PendudukResponse
	for _, penduduk := range penduduk {
		pendudukResponse = append(pendudukResponse, responses.PendudukResponseFromModel(&penduduk))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "OK",
		Message: "penduduk retrieved successfully",
		Data:    pendudukResponse,
	})

}

// get by id
func (h *PendudukHandler) GetPendudukByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid penduduk ID",
			Data:    nil,
		})
		return
	}

	penduduk, err := h.service.GetPendudukByID(uint(id))
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
		Message: "penduduk retrieved successfully",
		Data:    responses.PendudukResponseFromModel(penduduk),
	})
}

func (h *PendudukHandler) CountPenduduk(c *gin.Context) {
	penduduk, err := h.service.CountPenduduk()
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
		Message: "penduduk retrieved successfully",
		Data:    penduduk,
	})
}

// update
func (h *PendudukHandler) UpdatePenduduk(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid penduduk ID",
			Data:    nil,
		})
		return
	}

	var req requests.PendudukRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	penduduk, err := h.service.UpdatePenduduk(uint(id), &req)
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
		Message: "penduduk updated successfully",
		Data:    responses.PendudukResponseFromModel(penduduk),
	})
}

// delete
func (h *PendudukHandler) DeletePenduduk(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid penduduk ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeletePenduduk(uint(id))
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
		Message: "penduduk deleted successfully",
		Data:    nil,
	})
}
