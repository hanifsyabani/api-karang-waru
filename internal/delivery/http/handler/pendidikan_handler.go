package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PendidikanHandler struct {
	service usecase.PendidikanService
}

func NewPendidikanHandler(service usecase.PendidikanService) *PendidikanHandler {
	return &PendidikanHandler{service}
}

// ===============================
// 1. Lembaga Pendidikan
// ===============================
func (h *PendidikanHandler) CreateLembaga(c *gin.Context) {
	var req types.LembagaPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.CreateLembagaPendidikan(&req)
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
		Message: "Lembaga pendidikan created successfully",
		Data:    types.LembagaPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) GetAllLembaga(c *gin.Context) {
	data, err := h.service.GetAllLembagaPendidikan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var list []types.LembagaPendidikanResponse
	for _, d := range data {
		list = append(list, types.LembagaPendidikanFromModel(&d))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Lembaga pendidikan retrieved successfully",
		Data:    list,
	})
}

func (h *PendidikanHandler) GetLembagaByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.service.GetLembagaPendidikanByID(uint(id))
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
		Message: "Lembaga pendidikan retrieved successfully",
		Data:    types.LembagaPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) UpdateLembaga(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req types.LembagaPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.UpdateLembagaPendidikan(uint(id), &req)
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
		Message: "Lembaga pendidikan updated successfully",
		Data:    types.LembagaPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) DeleteLembaga(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteLembagaPendidikan(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Lembaga pendidikan deleted successfully",
		Data:    nil,
	})
}

// ===============================
// 2. Statistik Pendidikan
// ===============================
func (h *PendidikanHandler) CreateStatistik(c *gin.Context) {
	var req types.StatistikPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.CreateStatistikPendidikan(&req)
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
		Message: "Statistik created successfully",
		Data:    types.StatistikPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) GetAllStatistik(c *gin.Context) {
	data, err := h.service.GetAllStatistikPendidikan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var list []types.StatistikPendidikanResponse
	for _, d := range data {
		list = append(list, types.StatistikPendidikanFromModel(&d))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Statistik pendidikan retrieved successfully",
		Data:    list,
	})
}

func (h *PendidikanHandler) GetStatistikByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.service.GetStatistikPendidikanByID(uint(id))
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
		Message: "Statistik pendidikan retrieved successfully",
		Data:    types.StatistikPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) UpdateStatistik(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req types.StatistikPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.UpdateStatistikPendidikan(uint(id), &req)
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
		Message: "Statistik pendidikan updated successfully",
		Data:    types.StatistikPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) DeleteStatistik(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteStatistikPendidikan(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Statistik pendidikan deleted successfully",
		Data:    nil,
	})
}

// ===============================
// 3. Program Pendidikan
// ===============================
func (h *PendidikanHandler) CreateProgram(c *gin.Context) {
	var req types.ProgramPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.CreateProgramPendidikan(&req)
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
		Message: "Program created successfully",
		Data:    types.ProgramPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) GetAllProgram(c *gin.Context) {
	data, err := h.service.GetAllProgramPendidikan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var list []types.ProgramPendidikanResponse
	for _, d := range data {
		list = append(list, types.ProgramPendidikanFromModel(&d))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Program pendidikan retrieved successfully",
		Data:    list,
	})
}

func (h *PendidikanHandler) GetProgramByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.service.GetProgramPendidikanByID(uint(id))
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
		Message: "Program pendidikan retrieved successfully",
		Data:    types.ProgramPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) UpdateProgram(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req types.ProgramPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.UpdateProgramPendidikan(uint(id), &req)
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
		Message: "Program pendidikan updated successfully",
		Data:    types.ProgramPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) DeleteProgram(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteProgramPendidikan(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Program pendidikan deleted successfully",
		Data:    nil,
	})
}

// ===============================
// 4. Capaian Pendidikan
// ===============================
func (h *PendidikanHandler) CreateCapaian(c *gin.Context) {
	var req types.CapaianPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.CreateCapaianPendidikan(&req)
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
		Message: "Capaian created successfully",
		Data:    types.CapaianPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) GetAllCapaian(c *gin.Context) {
	data, err := h.service.GetAllCapaianPendidikan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var list []types.CapaianPendidikanResponse
	for _, d := range data {
		list = append(list, types.CapaianPendidikanFromModel(&d))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Capaian pendidikan retrieved successfully",
		Data:    list,
	})
}

func (h *PendidikanHandler) GetCapaianByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.service.GetCapaianPendidikanByID(uint(id))
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
		Message: "Capaian pendidikan retrieved successfully",
		Data:    types.CapaianPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) UpdateCapaian(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req types.CapaianPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.UpdateCapaianPendidikan(uint(id), &req)
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
		Message: "Capaian pendidikan updated successfully",
		Data:    types.CapaianPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) DeleteCapaian(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteCapaianPendidikan(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Capaian pendidikan deleted successfully",
		Data:    nil,
	})
}

// ===============================
// 5. Dokumentasi Pendidikan
// ===============================
func (h *PendidikanHandler) CreateDokumentasi(c *gin.Context) {
	var req types.DokumentasiPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.CreateDokumentasiPendidikan(&req)
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
		Message: "Dokumentasi created successfully",
		Data:    types.DokumentasiPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) GetAllDokumentasi(c *gin.Context) {
	data, err := h.service.GetAllDokumentasiPendidikan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var list []types.DokumentasiPendidikanResponse
	for _, d := range data {
		list = append(list, types.DokumentasiPendidikanFromModel(&d))
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Dokumentasi pendidikan retrieved successfully",
		Data:    list,
	})
}

func (h *PendidikanHandler) GetDokumentasiByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.service.GetDokumentasiPendidikanByID(uint(id))
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
		Message: "Dokumentasi pendidikan retrieved successfully",
		Data:    types.DokumentasiPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) UpdateDokumentasi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req types.DokumentasiPendidikanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.service.UpdateDokumentasiPendidikan(uint(id), &req)
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
		Message: "Dokumentasi pendidikan updated successfully",
		Data:    types.DokumentasiPendidikanFromModel(data),
	})
}

func (h *PendidikanHandler) DeleteDokumentasi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteDokumentasiPendidikan(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Code:    "OK",
		Message: "Dokumentasi pendidikan deleted successfully",
		Data:    nil,
	})
}
