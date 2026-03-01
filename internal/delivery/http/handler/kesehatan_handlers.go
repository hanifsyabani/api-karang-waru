package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KesehatanHandler struct {
	service usecase.KesehatanService
}

func NewKesehatanHandler(service usecase.KesehatanService) *KesehatanHandler {
	return &KesehatanHandler{service}
}

// ===============================
// CREATE
// ===============================
func (h *KesehatanHandler) CreateLayanan(c *gin.Context) {
	var req types.LayananKesehatanRequest

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

	c.JSON(http.StatusCreated, types.APIResponse{
		Code:    "CREATED",
		Message: "Layanan kesehatan created successfully",
		Data:    types.LayananKesehatanResponseFromModel(layanan),
	})
}

// ===============================
// GET ALL
// ===============================
func (h *KesehatanHandler) GetLayanan(c *gin.Context) {
	layanans, err := h.service.GetAllLayanan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var list []types.LayananKesehatanResponse
	for _, l := range layanans {
		list = append(list, types.LayananKesehatanResponseFromModel(&l))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Layanan kesehatan retrieved successfully",
		Data:    list,
	})
}

// ===============================
// GET BY ID
// ===============================
func (h *KesehatanHandler) GetLayananByID(c *gin.Context) {
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
		Message: "Layanan kesehatan retrieved successfully",
		Data:    types.LayananKesehatanResponseFromModel(layanan),
	})
}

// ===============================
// UPDATE
// ===============================
func (h *KesehatanHandler) UpdateLayanan(c *gin.Context) {
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

	var req types.LayananKesehatanRequest

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
		Message: "Layanan kesehatan updated successfully",
		Data:    types.LayananKesehatanResponseFromModel(layanan),
	})
}

// ===============================
// DELETE
// ===============================
func (h *KesehatanHandler) DeleteLayanan(c *gin.Context) {
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
		Message: "Layanan kesehatan deleted successfully",
		Data:    nil,
	})
}

func (h *KesehatanHandler) CreateFasilitas(c *gin.Context) {
	var req types.FasilitasKesehatanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	fasilitas, err := h.service.CreateFasilitas(&req)
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
		Message: "Fasilitas kesehatan created successfully",
		Data:    types.FasilitasKesehatanResponseFromModel(fasilitas),
	})
}

// ===============================
// GET ALL
// ===============================
func (h *KesehatanHandler) GetFasilitasKesehatan(c *gin.Context) {
	fasilitas, err := h.service.GetAllFasilitas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var list []types.FasilitasKesehatanResponse
	for _, l := range fasilitas {
		list = append(list, types.FasilitasKesehatanResponseFromModel(&l))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Fasilitas kesehatan retrieved successfully",
		Data:    list,
	})
}

// ===============================
// GET BY ID
// ===============================
func (h *KesehatanHandler) GetFasilitasKesehatanByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid fasilitas ID",
			Data:    nil,
		})
		return
	}

	fasilitas, err := h.service.GetFasilitasByID(uint(id))
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
		Message: "Fasilitas kesehatan retrieved successfully",
		Data:    types.FasilitasKesehatanResponseFromModel(fasilitas),
	})
}

// ===============================
// UPDATE
// ===============================
func (h *KesehatanHandler) UpdateFasilitasKesehatan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid fasilitas ID",
			Data:    nil,
		})
		return
	}

	var req types.FasilitasKesehatanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	fasilitas, err := h.service.UpdateFasilitas(uint(id), &req)
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
		Message: "Fasilitas kesehatan updated successfully",
		Data:    types.FasilitasKesehatanResponseFromModel(fasilitas),
	})
}

// ===============================
// DELETE
// ===============================
func (h *KesehatanHandler) DeleteFasilitasKesehatan(c *gin.Context) {
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

	err = h.service.DeleteFasilitas(uint(id))
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
		Message: "Fasilitas kesehatan deleted successfully",
		Data:    nil,
	})
}

