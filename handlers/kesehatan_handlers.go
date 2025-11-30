package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KesehatanHandler struct {
	service services.KesehatanService
}

func NewKesehatanHandler(service services.KesehatanService) *KesehatanHandler {
	return &KesehatanHandler{service}
}

// ===============================
// CREATE
// ===============================
func (h *KesehatanHandler) CreateLayanan(c *gin.Context) {
	var req requests.LayananKesehatanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	layanan, err := h.service.CreateLayanan(&req)
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
		Message: "Layanan kesehatan created successfully",
		Data:    responses.LayananKesehatanResponseFromModel(layanan),
	})
}

// ===============================
// GET ALL
// ===============================
func (h *KesehatanHandler) GetLayanan(c *gin.Context) {
	layanans, err := h.service.GetAllLayanan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var list []responses.LayananKesehatanResponse
	for _, l := range layanans {
		list = append(list, responses.LayananKesehatanResponseFromModel(&l))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
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
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid layanan ID",
			Data:    nil,
		})
		return
	}

	layanan, err := h.service.GetLayananByID(uint(id))
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
		Message: "Layanan kesehatan retrieved successfully",
		Data:    responses.LayananKesehatanResponseFromModel(layanan),
	})
}

// ===============================
// UPDATE
// ===============================
func (h *KesehatanHandler) UpdateLayanan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid layanan ID",
			Data:    nil,
		})
		return
	}

	var req requests.LayananKesehatanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	layanan, err := h.service.UpdateLayanan(uint(id), &req)
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
		Message: "Layanan kesehatan updated successfully",
		Data:    responses.LayananKesehatanResponseFromModel(layanan),
	})
}

// ===============================
// DELETE
// ===============================
func (h *KesehatanHandler) DeleteLayanan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid layanan ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeleteLayanan(uint(id))
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
		Message: "Layanan kesehatan deleted successfully",
		Data:    nil,
	})
}




func (h *KesehatanHandler) CreateFasilitas(c *gin.Context) {
	var req requests.FasilitasKesehatanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	fasilitas, err := h.service.CreateFasilitas(&req)
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
		Message: "Fasilitas kesehatan created successfully",
		Data:    responses.FasilitasKesehatanResponseFromModel(fasilitas),
	})
}

// ===============================
// GET ALL
// ===============================
func (h *KesehatanHandler) GetFasilitasKesehatan(c *gin.Context) {
	fasilitas, err := h.service.GetAllFasilitas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var list []responses.FasilitasKesehatanResponse
	for _, l := range fasilitas {
		list = append(list, responses.FasilitasKesehatanResponseFromModel(&l))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
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
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid fasilitas ID",
			Data:    nil,
		})
		return
	}

	fasilitas, err := h.service.GetFasilitasByID(uint(id))
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
		Message: "Fasilitas kesehatan retrieved successfully",
		Data:    responses.FasilitasKesehatanResponseFromModel(fasilitas),
	})
}

// ===============================
// UPDATE
// ===============================
func (h *KesehatanHandler) UpdateFasilitasKesehatan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid fasilitas ID",
			Data:    nil,
		})
		return
	}

	var req requests.FasilitasKesehatanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	fasilitas, err := h.service.UpdateFasilitas(uint(id), &req)
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
		Message: "Fasilitas kesehatan updated successfully",
		Data:    responses.FasilitasKesehatanResponseFromModel(fasilitas),
	})
}

// ===============================
// DELETE
// ===============================
func (h *KesehatanHandler) DeleteFasilitasKesehatan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid layanan ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeleteFasilitas(uint(id))
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
		Message: "Fasilitas kesehatan deleted successfully",
		Data:    nil,
	})
}
