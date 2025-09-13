package handlers

import (
	"api-karang-waru/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainHandler struct{}

func NewMainHandler() *MainHandler {
	return &MainHandler{}
}

// MainHandler responds with a message indicating that the API Contact Form is running.
func (h *MainHandler) MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "API Contact Form is running.",
	})
}