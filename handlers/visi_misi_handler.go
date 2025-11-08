package handlers

import (
	"api-karang-waru/requests"
	"api-karang-waru/responses"
	"api-karang-waru/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VisiMisiHandler struct {
	service services.VisiMisiService
}

func NewVisiMisiHandler(service services.VisiMisiService) *VisiMisiHandler {
	return &VisiMisiHandler{service}
}


func (h *VisiMisiHandler) CreateVisiMisi(c *gin.Context) {

	var req requests.VisiMisiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	visiMisi, err := h.service.CreateVisiMisi(&req)

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
		Message: "Visi Misi created successfully",
		Data:    responses.VisiMisiResponseFromModel(visiMisi),
	})
}

func (h *VisiMisiHandler) GetVisiMisi(c *gin.Context) {
	visiMisi, err := h.service.GetVisiMisi()

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
		Message: "Visi Misi retrieved successfully",
		Data:    responses.VisiMisiResponseFromModel(&visiMisi),
	})

}

func (h *VisiMisiHandler) UpdateVisiMisi(c *gin.Context) {

	var req requests.VisiMisiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	visiMisi, err := h.service.UpdateVisiMisi(&req)
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
		Message: "Visi Misi updated successfully",
		Data:    responses.VisiMisiResponseFromModel(visiMisi),
	})
}

// delete
func (h *VisiMisiHandler) DeleteVisiMisi(c *gin.Context) {

	err := h.service.DeleteVisiMisi()
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
		Message: "Visi Misi deleted successfully",
		Data:    nil,
	})
}