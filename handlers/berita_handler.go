package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BeritaHandler struct {
	service services.BeritaService
}

func NewBeritaHandler(service services.BeritaService) *BeritaHandler {
	return &BeritaHandler{service}
}

func (h *BeritaHandler) CreateBerita(c *gin.Context) {
	// BeritaRequest → fokus pada data yang dikirim berita.
	// models.Berita → fokus pada struktur tabel di database.

	var req requests.BeritaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	berita, err := h.service.CreateBerita(&req)

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
		Message: "Berita created successfully",
		Data:    responses.BeritaResponseFromModel(berita),
	})
}

// get all
func (h *BeritaHandler) GetBerita(c *gin.Context) {
	berita, err := h.service.GetAllBerita()

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var beritaResponses []responses.BeritaResponse
	for _, berita := range berita {
		beritaResponses = append(beritaResponses, responses.BeritaResponseFromModel(&berita))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "OK",
		Message: "Berita retrieved successfully",
		Data:    beritaResponses,
	})

}

// get by id
func (h *BeritaHandler) GetBeritaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid berita ID",
			Data:    nil,
		})
		return
	}

	berita, err := h.service.GetBeritaByID(uint(id))
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
		Message: "berita retrieved successfully",
		Data:    responses.BeritaResponseFromModel(berita),
	})
}

func (h *BeritaHandler) GetBeritaBySlug(c *gin.Context) {
	slug := c.Param("slug")
	berita, err := h.service.GetBeritaBySlug(slug)
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
		Message: "berita retrieved successfully",
		Data:    responses.BeritaResponseFromModel(berita),
	})
}

// update
func (h *BeritaHandler) UpdateBerita(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid berita ID",
			Data:    nil,
		})
		return
	}

	var req requests.BeritaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	berita, err := h.service.UpdateBerita(uint(id), &req)
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
		Data:    responses.BeritaResponseFromModel(berita),
	})
}

// delete
func (h *BeritaHandler) DeleteBerita(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid berita ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeleteBerita(uint(id))
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
		Message: "Berita deleted successfully",
		Data:    nil,
	})
}
