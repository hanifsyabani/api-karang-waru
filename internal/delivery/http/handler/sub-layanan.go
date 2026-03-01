package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubLayananHandler struct {
	service usecase.SubLayananService
}

func NewSubLayananHandler(service usecase.SubLayananService) *SubLayananHandler {
	return &SubLayananHandler{service}
}

// CreateSubLayanan godoc
// @Summary Create sub layanan
// @Description Membuat sub layanan baru
// @Tags Sub Layanan
// @Accept json
// @Produce json
// @Param body body types.SubLayananRequest true "Sub Layanan Request"
// @Success 201 {object} types.APIResponse
// @Failure 400 {object} types.APIResponse
// @Failure 401 {object} types.APIResponse
// @Security BearerAuth
// @Router /sub-service [post]
func (h *SubLayananHandler) CreateSubLayanan(c *gin.Context) {
	var req types.SubLayananRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	subLayanan, err := h.service.CreateSubLayanan(&req)

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
		Message: "Sub layanan created successfully",
		Data:    types.SubLayananResponseFromModel(subLayanan),
	})
}

// GetAllSubLayanan godoc
// @Summary Get all sub layanan
// @Description Ambil daftar sub layanan (pagination & filter)
// @Tags Sub Layanan
// @Produce json
// @Param search query string false "Search keyword"
// @Param layanan_desa_id query int false "Layanan Desa ID"
// @Param page query int false "Page number"
// @Param limit query int false "Limit per page"
// @Param sort_by query string false "Sort by field"
// @Param sort_order query string false "asc / desc"
// @Success 200 {object} types.APIResponse
// @Security BearerAuth
// @Router /sub-service [get]
func (h *SubLayananHandler) GetAllSubLayanan(c *gin.Context) {
	search := c.Query("search")
	layananDesaIDParam := c.Query("layanan_desa_id")
	pageParam, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limitParam, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sortBy := c.DefaultQuery("sort_by", "created_at")
	sortOrder := c.DefaultQuery("sort_order", "desc")

	var layananDesaID uint
	if layananDesaIDParam != "" {
		id, _ := strconv.Atoi(layananDesaIDParam)
		layananDesaID = uint(id)
	}

	subLayanan, err := h.service.GetAllSubLayanan(layananDesaID, search, pageParam, limitParam, sortBy, sortOrder)

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var subLayananResponse []types.SubLayananResponse
	for _, sub := range subLayanan {
		subLayananResponse = append(subLayananResponse, types.SubLayananResponseFromModel(&sub))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Sub layanan retrieved successfully",
		Data:    subLayananResponse,
	})
}

func (h *SubLayananHandler) GetSubLayananByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid sub layanan ID",
			Data:    nil,
		})
		return
	}

	subLayanan, err := h.service.GetSubLayananByID(uint(id))
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
		Message: "Sub layanan retrieved successfully",
		Data:    types.SubLayananResponseFromModel(subLayanan),
	})
}

func (h *SubLayananHandler) GetSubLayananByLayananID(c *gin.Context) {
	layananIDParam := c.Param("layanan_id")
	layananID, err := strconv.Atoi(layananIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid layanan ID",
			Data:    nil,
		})
		return
	}

	subLayanan, err := h.service.GetSubLayananByLayananID(uint(layananID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var subLayananResponse []types.SubLayananResponse
	for _, sub := range subLayanan {
		subLayananResponse = append(subLayananResponse, types.SubLayananResponseFromModel(&sub))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Sub layanan retrieved successfully",
		Data:    subLayananResponse,
	})
}

// UpdateSubLayanan godoc
// @Summary Update sub layanan
// @Description Update sub layanan berdasarkan ID
// @Tags Sub Layanan
// @Accept json
// @Produce json
// @Param id path int true "Sub Layanan ID"
// @Param body body types.SubLayananRequest true "Update request"
// @Success 200 {object} types.APIResponse
// @Failure 400 {object} types.APIResponse
// @Failure 404 {object} types.APIResponse
// @Security BearerAuth
// @Router /sub-service/{id} [put]
func (h *SubLayananHandler) UpdateSubLayanan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid sub layanan ID",
			Data:    nil,
		})
		return
	}

	var req types.SubLayananRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	subLayanan, err := h.service.UpdateSubLayanan(uint(id), &req)
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
		Message: "Sub layanan updated successfully",
		Data:    types.SubLayananResponseFromModel(subLayanan),
	})
}

func (h *SubLayananHandler) DeleteSubLayanan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid sub layanan ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeleteSubLayanan(uint(id))
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
		Message: "Sub layanan deleted successfully",
		Data:    nil,
	})
}

// ==================== PENGAJUAN LAYANAN HANDLER ====================

type PengajuanLayananHandler struct {
	service usecase.PengajuanLayananService
}

func NewPengajuanLayananHandler(service usecase.PengajuanLayananService) *PengajuanLayananHandler {
	return &PengajuanLayananHandler{service}
}

func (h *PengajuanLayananHandler) CreatePengajuan(c *gin.Context) {
	var req types.PengajuanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	pengajuan, err := h.service.CreatePengajuan(&req)

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
		Message: "Pengajuan created successfully",
		Data:    types.PengajuanResponseFromModel(pengajuan),
	})
}

func (h *PengajuanLayananHandler) GetAllPengajuan(c *gin.Context) {
	search := c.Query("search")
	status := c.Query("status")
	nik := c.Query("nik")
	layananDesaIDParam := c.Query("layanan_desa_id")
	pageParam, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limitParam, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sortBy := c.DefaultQuery("sort_by", "created_at")
	sortOrder := c.DefaultQuery("sort_order", "desc")

	var layananDesaID uint
	if layananDesaIDParam != "" {
		id, _ := strconv.Atoi(layananDesaIDParam)
		layananDesaID = uint(id)
	}

	pengajuan, err := h.service.GetAllPengajuan(search, status, nik, layananDesaID, pageParam, limitParam, sortBy, sortOrder)

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var pengajuanResponse []types.PengajuanResponse
	for _, p := range pengajuan {
		pengajuanResponse = append(pengajuanResponse, types.PengajuanResponseFromModel(&p))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Pengajuan retrieved successfully",
		Data:    pengajuanResponse,
	})
}

func (h *PengajuanLayananHandler) GetPengajuanByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid pengajuan ID",
			Data:    nil,
		})
		return
	}

	pengajuan, err := h.service.GetPengajuanByID(uint(id))
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
		Message: "Pengajuan retrieved successfully",
		Data:    types.PengajuanResponseFromModel(pengajuan),
	})
}

func (h *PengajuanLayananHandler) GetPengajuanByNIK(c *gin.Context) {
	nik := c.Param("nik")

	pengajuan, err := h.service.GetPengajuanByNIK(nik)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var pengajuanResponse []types.PengajuanResponse
	for _, p := range pengajuan {
		pengajuanResponse = append(pengajuanResponse, types.PengajuanResponseFromModel(&p))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Pengajuan retrieved successfully",
		Data:    pengajuanResponse,
	})
}

func (h *PengajuanLayananHandler) GetPengajuanByNomorSurat(c *gin.Context) {
	nomorSurat := c.Param("nomor_surat")

	pengajuan, err := h.service.GetPengajuanByNomorSurat(nomorSurat)
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
		Message: "Pengajuan retrieved successfully",
		Data:    types.PengajuanResponseFromModel(pengajuan),
	})
}

func (h *PengajuanLayananHandler) UpdatePengajuan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid pengajuan ID",
			Data:    nil,
		})
		return
	}

	var req types.PengajuanUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	pengajuan, err := h.service.UpdatePengajuan(uint(id), &req)
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
		Message: "Pengajuan updated successfully",
		Data:    types.PengajuanResponseFromModel(pengajuan),
	})
}

func (h *PengajuanLayananHandler) UpdateStatusPengajuan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid pengajuan ID",
			Data:    nil,
		})
		return
	}

	var req types.UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// TODO: Get user ID from JWT token/session
	// For now, using dummy user ID
	updatedBy := uint(1)

	pengajuan, err := h.service.UpdateStatusPengajuan(uint(id), &req, updatedBy)
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
		Message: "Status pengajuan updated successfully",
		Data:    types.PengajuanResponseFromModel(pengajuan),
	})
}

func (h *PengajuanLayananHandler) ApprovePengajuan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid pengajuan ID",
			Data:    nil,
		})
		return
	}

	var req types.ApproveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// TODO: Get user ID from JWT token/session
	approvedBy := uint(1)

	pengajuan, err := h.service.ApprovePengajuan(uint(id), &req, approvedBy)
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
		Message: "Pengajuan approved successfully",
		Data:    types.PengajuanResponseFromModel(pengajuan),
	})
}

func (h *PengajuanLayananHandler) RejectPengajuan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid pengajuan ID",
			Data:    nil,
		})
		return
	}

	var req types.RejectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// TODO: Get user ID from JWT token/session
	rejectedBy := uint(1)

	pengajuan, err := h.service.RejectPengajuan(uint(id), &req, rejectedBy)
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
		Message: "Pengajuan rejected successfully",
		Data:    types.PengajuanResponseFromModel(pengajuan),
	})
}

func (h *PengajuanLayananHandler) DeletePengajuan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid pengajuan ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeletePengajuan(uint(id))
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
		Message: "Pengajuan deleted successfully",
		Data:    nil,
	})
}

func (h *PengajuanLayananHandler) GetStatisticsByStatus(c *gin.Context) {
	statistics, err := h.service.GetStatisticsByStatus()
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
		Message: "Statistics retrieved successfully",
		Data:    statistics,
	})
}

// ==================== RIWAYAT PENGAJUAN HANDLER ====================

type RiwayatPengajuanHandler struct {
	service usecase.RiwayatPengajuanService
}

func NewRiwayatPengajuanHandler(service usecase.RiwayatPengajuanService) *RiwayatPengajuanHandler {
	return &RiwayatPengajuanHandler{service}
}

func (h *RiwayatPengajuanHandler) GetRiwayatByPengajuanID(c *gin.Context) {
	pengajuanIDParam := c.Param("pengajuan_id")
	pengajuanID, err := strconv.Atoi(pengajuanIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid pengajuan ID",
			Data:    nil,
		})
		return
	}

	riwayat, err := h.service.GetRiwayatByPengajuanID(uint(pengajuanID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var riwayatResponse []types.RiwayatPengajuanResponse
	for _, r := range riwayat {
		riwayatResponse = append(riwayatResponse, types.RiwayatPengajuanResponseFromModel(&r))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Riwayat pengajuan retrieved successfully",
		Data:    riwayatResponse,
	})
}

func (h *RiwayatPengajuanHandler) GetAllRiwayat(c *gin.Context) {
	pengajuanIDParam := c.Query("pengajuan_id")
	pageParam, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limitParam, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var pengajuanID uint
	if pengajuanIDParam != "" {
		id, _ := strconv.Atoi(pengajuanIDParam)
		pengajuanID = uint(id)
	}

	riwayat, err := h.service.GetAllRiwayat(pengajuanID, pageParam, limitParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var riwayatResponse []types.RiwayatPengajuanResponse
	for _, r := range riwayat {
		riwayatResponse = append(riwayatResponse, types.RiwayatPengajuanResponseFromModel(&r))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Riwayat pengajuan retrieved successfully",
		Data:    riwayatResponse,
	})
}
