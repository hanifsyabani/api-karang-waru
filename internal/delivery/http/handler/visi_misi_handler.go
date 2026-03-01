package handler

import (
	"api-karang-waru/pkg/types"
	"api-karang-waru/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VisiMisiHandler struct {
	service usecase.VisiMisiService
}

func NewVisiMisiHandler(service usecase.VisiMisiService) *VisiMisiHandler {
	return &VisiMisiHandler{service}
}

func (h *VisiMisiHandler) CreateVisiMisi(c *gin.Context) {

	var req types.VisiMisiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	visiMisi, err := h.service.CreateVisiMisi(&req)

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
		Message: "Visi Misi created successfully",
		Data:    types.VisiMisiResponseFromModel(visiMisi),
	})
}

func (h *VisiMisiHandler) GetVisiMisi(c *gin.Context) {
	visiMisi, err := h.service.GetVisiMisi()

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
		Message: "Visi Misi retrieved successfully",
		Data:    types.VisiMisiResponseFromModel(&visiMisi),
	})

}

func (h *VisiMisiHandler) UpdateVisiMisi(c *gin.Context) {

	var req types.VisiMisiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	visiMisi, err := h.service.UpdateVisiMisi(&req)
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
		Message: "Visi Misi updated successfully",
		Data:    types.VisiMisiResponseFromModel(visiMisi),
	})
}

// delete
func (h *VisiMisiHandler) DeleteVisiMisi(c *gin.Context) {

	err := h.service.DeleteVisiMisi()
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
		Message: "Visi Misi deleted successfully",
		Data:    nil,
	})
}
