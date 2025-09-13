package handlers

import (
	"api-karang-waru/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles HTTP requests related to health checks.
type HealthHandler struct{}

// NewHealthHandler creates a new instance of HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck responds with a simple message indicating that the API is running.
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "API is running.",
	})
}