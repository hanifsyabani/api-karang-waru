package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PendudukHandler struct {
	service usecase.PendudukService
}

func NewPendudukHandler(service usecase.PendudukService) *PendudukHandler {
	return &PendudukHandler{service}
}

func (h *PendudukHandler) CreatePenduduk(c *gin.Context) {
	var req types.PendudukRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	penduduk, err := h.service.CreatePenduduk(&req)

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
		Message: "penduduk created successfully",
		Data:    types.PendudukResponseFromModel(penduduk),
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
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var pendudukResponse []types.PendudukResponse
	for _, penduduk := range penduduk {
		pendudukResponse = append(pendudukResponse, types.PendudukResponseFromModel(&penduduk))
	}

	c.JSON(http.StatusOK, types.APIResponse{
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
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid penduduk ID",
			Data:    nil,
		})
		return
	}

	penduduk, err := h.service.GetPendudukByID(uint(id))
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
		Message: "penduduk retrieved successfully",
		Data:    types.PendudukResponseFromModel(penduduk),
	})
}

func (h *PendudukHandler) CountPenduduk(c *gin.Context) {
	total, lakiLaki, perempuan, kartuKeluarga, err := h.service.CountPenduduk()
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
		Message: "penduduk retrieved successfully",
		Data:    map[string]interface{}{
			"total":          total,
			"male":      lakiLaki,
			"female":      perempuan,
			"family_card_count": kartuKeluarga,
		},
	})
}

// update
func (h *PendudukHandler) UpdatePenduduk(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid penduduk ID",
			Data:    nil,
		})
		return
	}

	var req types.PendudukRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	penduduk, err := h.service.UpdatePenduduk(uint(id), &req)
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
		Message: "penduduk updated successfully",
		Data:    types.PendudukResponseFromModel(penduduk),
	})
}

// delete
func (h *PendudukHandler) DeletePenduduk(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid penduduk ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeletePenduduk(uint(id))
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
		Message: "penduduk deleted successfully",
		Data:    nil,
	})
}

