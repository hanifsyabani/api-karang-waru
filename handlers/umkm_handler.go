package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UmkmHandler struct {
	service services.UmkmService
}

func NewUmkmHandler(service services.UmkmService) *UmkmHandler {
	return &UmkmHandler{service}
}

func (h *UmkmHandler) CreateUmkm(c *gin.Context) {
	// UmkmRequest → fokus pada data yang dikirim umkm.
	// models.Umkm → fokus pada struktur tabel di database.

	var req requests.UmkmRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	umkm, err := h.service.CreateUmkm(&req)

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
		Message: "Umkm created successfully",
		Data:    responses.UmkmResponseFromModel(umkm),
	})
}

// get all
func (h *UmkmHandler) GetAllUmkm(c *gin.Context) {
	umkm, err := h.service.GetAllUmkm()

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var UmkmResponse []responses.UmkmResponse
	for _, umkm := range umkm {
		UmkmResponse = append(UmkmResponse, responses.UmkmResponseFromModel(&umkm))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "OK",
		Message: "Umkm retrieved successfully",
		Data:    UmkmResponse,
	})

}

// get by id
func (h *UmkmHandler) GetUmkmByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid umkm ID",
			Data:    nil,
		})
		return
	}

	umkm, err := h.service.GetUmkmByID(uint(id))
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
		Message: "umkm retrieved successfully",
		Data:    responses.UmkmResponseFromModel(umkm),
	})
}

func (h *UmkmHandler) GetUmkmBySlug(c *gin.Context) {
	slug := c.Param("slug")
	umkm, err := h.service.GetUmkmBySlug(slug)
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
		Message: "umkm retrieved successfully",
		Data:    responses.UmkmResponseFromModel(umkm),
	})
}

// update
func (h *UmkmHandler) UpdateUmkm(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid umkm ID",
			Data:    nil,
		})
		return
	}

	var req requests.UmkmRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	umkm, err := h.service.UpdateUmkm(uint(id), &req)
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
		Message: "Umkm updated successfully",
		Data:    responses.UmkmResponseFromModel(umkm),
	})
}

// delete
func (h *UmkmHandler) DeleteUmkm(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid umkm ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeleteUmkm(uint(id))
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
		Message: "Umkm deleted successfully",
		Data:    nil,
	})
}
