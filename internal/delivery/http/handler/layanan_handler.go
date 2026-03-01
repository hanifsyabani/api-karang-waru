package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LayananHandler struct {
	service usecase.LayananService
}

func NewLayananHandler(service usecase.LayananService) *LayananHandler {
	return &LayananHandler{service}
}

func (h *LayananHandler) CreateLayanan(c *gin.Context) {
	var req types.LayananRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	layanan, err := h.service.CreateLayanan(&req)

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
		Message: "Layanan created successfully",
		Data:    types.LayananResponseFromModel(layanan),
	})
}

// get all
func (h *LayananHandler) GetAllLayanan(c *gin.Context) {

	search := c.Query("search")
	pageParam, _ := strconv.Atoi(c.DefaultQuery("page", "1")) 
	limitParam,_ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sortBy := c.DefaultQuery("sort_by", "created_at")
	sortOrder := c.DefaultQuery("sort_order", "desc")
	layanan, err := h.service.GetAllLayanan(search, pageParam, limitParam, sortBy, sortOrder)

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var LayananResponse []types.LayananResponse
	for _, layanan := range layanan {
		LayananResponse = append(LayananResponse, types.LayananResponseFromModel(&layanan))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Layanan retrieved successfully",
		Data:    LayananResponse,
	})

}

// get by id
func (h *LayananHandler) GetLayananByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid layanan ID",
			Data:    nil,
		})
		return
	}

	layanan, err := h.service.GetLayananByID(uint(id))
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
		Message: "layanan retrieved successfully",
		Data:    types.LayananResponseFromModel(layanan),
	})
}

func (h *LayananHandler) GetLayananBySlug(c *gin.Context) {
	slug := c.Param("slug")
	layanan, err := h.service.GetLayananBySlug(slug)
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
		Message: "layanan retrieved successfully",
		Data:    types.LayananResponseFromModel(layanan),
	})
}

// update
func (h *LayananHandler) UpdateLayanan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid layanan ID",
			Data:    nil,
		})
		return
	}

	var req types.LayananRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	layanan, err := h.service.UpdateLayanan(uint(id), &req)
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
		Message: "Layanan updated successfully",
		Data:    types.LayananResponseFromModel(layanan),
	})
}

// delete
func (h *LayananHandler) DeleteLayanan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid layanan ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeleteLayanan(uint(id))
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
		Message: "Layanan deleted successfully",
		Data:    nil,
	})
}

